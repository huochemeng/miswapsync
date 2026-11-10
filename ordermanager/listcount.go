package ordermanager

import (
	"MiSwapSync/stores/gdb"
	"context"
	"fmt"
	"github.com/pkg/errors"
	"log/slog"
	"strings"
	"time"
)

type CollectionListed struct {
	CollectionAddress string `json:"collection_address"`
	ListCount         int    `json:"list_count"`
}

func (om *OrderManager) listCountProcess(ctx context.Context) {
	// 启动时全量统计并缓存
	collectionsListed, err := om.countCollectionListed(ctx, []string{})
	if err != nil {
		slog.ErrorContext(ctx, "failed to count all collection listed on startup", "error", err)
	} else if err := om.cacheCollectionListCount(ctx, collectionsListed); err != nil {
		slog.ErrorContext(ctx, "failed to cache all collection listed count on startup", "error", err)
	}
	// ↑ 服务启动时立即执行一次全量查询，将所有 Collection 的有效挂单数写入 Redis。
	//   传入空切片表示不限制地址，即统计全部。若失败仅记录日志，不中断后续流程。

	// 记录待更新的集合地址（去重）
	pending := make(map[string]bool)
	// ↑ 用 map 做去重缓冲区：同一分钟内收到多次相同地址的变更通知时，只保留一个，
	//   避免对同一个 Collection 重复执行 SQL 查询。

	ticker := time.NewTicker(60 * time.Second)
	defer ticker.Stop()
	// ↑ 创建一个每 60 秒触发一次的定时器。ticker.C 是一个 <-chan time.Time 类型的通道，
	//   每隔 60 秒会向该通道发送一个当前时间值。defer Stop() 确保协程退出时释放定时器资源，
	//   防止 goroutine 泄漏。

	for {
		select {
		case <-ctx.Done():
			slog.InfoContext(ctx, "listCountProcess cancelled", "error", ctx.Err())
			return
			// ↑ 监听上下文取消信号。当父级调用 cancel() 或超时触发时，ctx.Done() 通道关闭，
			//   select 命中此分支，记录退出原因后安全返回。ctx.Err() 返回取消的具体原因
			//   （context.Canceled 或 context.DeadlineExceeded）。

		case addr := <-om.collectionListedCh:
			pending[strings.ToLower(addr)] = true
			// ↑ 监听上游（如 floorPriceProcess）发来的集合地址变更通知。
			//   收到地址后统一转小写存入 pending map，实现去重缓冲。
			//   此分支不会阻塞：如果通道为空，select 会继续等待其他分支。

		case <-ticker.C:
			if len(pending) == 0 {
				continue
			}
			// ↑ 每 60 秒触发一次。ticker.C 是定时器暴露的只读通道，每次到点时会产出一个 time.Time 值。
			//   这里用 <-ticker.C 仅消费信号、丢弃具体值（因为不需要时间戳本身）。
			//   若 pending 为空说明这一分钟内没有任何变更通知，直接跳过，避免无意义的空查询。

			// 提取待更新地址并清空
			addrs := make([]string, 0, len(pending))
			for a := range pending {
				addrs = append(addrs, a)
			}
			pending = make(map[string]bool)
			// ↑ 将 map 中的地址收集为切片，供批量 SQL 查询使用；
			//   然后立即重建空 map，开始下一轮缓冲。先提取再清空，保证本次处理的地址不会丢失。

			// 增量统计并缓存
			result, err := om.countCollectionListed(ctx, addrs)
			if err != nil {
				slog.ErrorContext(ctx, "failed to count collection listed for pending collections",
					"count", len(addrs), "error", err)
				continue
			}
			// ↑ 对待更新地址执行批量 COUNT(DISTINCT token_id) 查询。
			//   若查询失败，记录日志后 continue 进入下一次 select 循环，不退出协程。
			//   注意：失败的这批地址已被清空，不会自动重试，需等下次变更通知重新入队。

			if err := om.cacheCollectionListCount(ctx, result); err != nil {
				slog.ErrorContext(ctx, "failed to cache collection listed count for pending collections",
					"count", len(addrs), "error", err)
			}
			// ↑ 将查询结果逐条写入 Redis。即使部分 key 写入失败，也仅记录日志，
			//   不影响其他 key 的写入和后续循环。
		}
	}
}

// countCollectionListed 统计指定集合的有效挂单数，cs 为空时统计全部
func (om *OrderManager) countCollectionListed(ctx context.Context, cs []string) ([]CollectionListed, error) {
	itemTable := gdb.GetMultiProjectItemTableName(om.project, om.chain)
	orderTable := gdb.GetMultiProjectOrderTableName(om.project, om.chain)

	var query string
	var args []any

	if len(cs) > 0 {
		query = fmt.Sprintf(`
			SELECT co.collection_address, COUNT(DISTINCT co.token_id) AS list_count
			FROM %s AS ci
			JOIN %s AS co ON co.collection_address = ci.collection_address AND co.token_id = ci.token_id
			WHERE co.collection_address IN ?
			  AND co.order_type = 1
			  AND co.order_status = 0
			  AND co.maker = ci.owner
			  AND (ci.is_opensea_banned, co.marketplace_id) != (true, 1)
			GROUP BY co.collection_address`, itemTable, orderTable)
		args = []any{cs}
	} else {
		query = fmt.Sprintf(`
			SELECT co.collection_address, COUNT(DISTINCT co.token_id) AS list_count
			FROM %s AS ci
			JOIN %s AS co ON co.collection_address = ci.collection_address AND co.token_id = ci.token_id
			WHERE co.order_type = 1
			  AND co.order_status = 0
			  AND co.maker = ci.owner
			  AND (ci.is_opensea_banned, co.marketplace_id) != (true, 1)
			GROUP BY co.collection_address`, itemTable, orderTable)
	}

	var result []CollectionListed
	if err := om.DB.WithContext(ctx).Raw(query, args...).Scan(&result).Error; err != nil {
		return nil, errors.Wrap(err, "countCollectionListed query failed")
	}
	return result, nil
}

func (om *OrderManager) cacheCollectionListCount(ctx context.Context, collectionsListed []CollectionListed) error {
	for _, c := range collectionsListed {
		key := GenCollectionListedKey(om.chain, c.CollectionAddress)
		if err := om.Xkv.SetInt(key, c.ListCount); err != nil {
			slog.WarnContext(ctx, "failed to set collection listed count in redis",
				"key", key, "count", c.ListCount, "error", err)
			// 继续处理其余条目，不因单个失败中断
		}
	}
	return nil
}

func GenCollectionListedKey(chain, address string) string {
	return fmt.Sprintf("cache:ms:%s:collection:listed:%s", strings.ToLower(chain), strings.ToLower(address))
}
