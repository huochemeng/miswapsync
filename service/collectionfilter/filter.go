package collectionfilter

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"strings"
	"sync"
)

type Filter struct {
	//原写法：ctx定义在了结构体，这种是错误的写法。在高并发下，ctx就会错乱。推荐做法：ctx不要定义在结构体，而是需要的函数以参数形式添加
	//ctx context.Context
	db      *gorm.DB
	chain   string
	set     map[string]bool
	lock    *sync.RWMutex
	project string
}

func New(db *gorm.DB, chain string, project string) *Filter {
	return &Filter{
		db:      db,
		chain:   chain,
		set:     make(map[string]bool),
		lock:    &sync.RWMutex{},
		project: project,
	}
}

func (f *Filter) Add(element string) {
	f.lock.Lock()
	defer f.lock.Unlock()
	f.set[strings.ToLower(element)] = true
}

func (f *Filter) Remove(element string) {
	f.lock.Lock()
	defer f.lock.Unlock()
	delete(f.set, strings.ToLower(element))
}

func (f *Filter) Contains(element string) bool {
	f.lock.RLock()         // Acquire the lock for reading
	defer f.lock.RUnlock() // Defer unlocking until function return
	_, exists := f.set[strings.ToLower(element)]
	return exists
}

func (f *Filter) PreloadCollections(ctx context.Context) error {
	var addresses []string

	//query the addresses directly from database
	err := f.db.WithContext(ctx).
		Table("db_sync").
		Select("address").
		Where("floor_price_status = ?", 1).
		Scan(&addresses).Error
	if err != nil {
		return errors.Wrap(err, "failed to query collections from database")
	}
	// 原写法：使用Add频繁获取/释放锁带来大开销
	//for _, address := range addresses {
	//	f.Add(address)
	//}

	// 避免在循环里频繁获取/释放锁带来的开销
	f.lock.Lock()
	defer f.lock.Unlock()
	for _, address := range addresses {
		f.set[strings.ToLower(address)] = true
	}

	return nil
}
