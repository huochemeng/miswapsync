# MiSwapSync
> Chain-Event Indexer & Orderbook Sync Service

## 📖 项目介绍

本项目是一个高性能的链上事件监听与数据索引服务，专为去中心化订单簿（如 MiSwapOrderBook 等 NFT/DEX 合约）设计。服务通过实时监听链上合约事件（如订单创建、取消 `LogCancel`、地板价变更等），解析并持久化关键数据，为后端的订单撮合引擎和前端数据展示提供可靠的数据源支持。

项目采用 Go 语言开发，基于 Cobra 构建命令行工具，支持多链、多项目隔离部署，并内置完善的后台维护机制与优雅关闭策略，确保在复杂网络环境下的数据一致性与服务稳定性。

---

## ✨ 核心特性

- **实时事件监听**：基于 `go-ethereum` 订阅或轮询链上日志，精准解析 `indexed` 参数（如 `OrderKey`, `Maker`）与 Data 负载。
- **多项目/多链支持**：通过动态表名和配置隔离，单套代码支持多链和多项目的并行索引。
- **高并发与协程模型**：使用 Goroutine 并行处理区块解析、数据清洗与数据库批量写入（Batch Upsert），最大化吞吐量。
- **后台维护机制 (Upkeep)**：内置定时清理过期数据、周期性校准集合地板价（Floor Price）等后台循环任务。
- **优雅关闭 (Graceful Shutdown)**：全链路基于 `context.Context` 传播取消信号，确保在接收到系统中断信号时，协程能安全退出、数据库连接正常释放，避免数据截断或脏写。
- **批量 Upsert 优化**：针对高频事件，采用 MySQL `ON DUPLICATE KEY UPDATE` 结合参数化占位符批量写入，兼顾性能与防 SQL 注入。

---

## 🛠 技术栈

| 组件 | 技术选型 | 说明 |
| :--- | :--- | :--- |
| **语言** | Go 1.20+ | 核心开发语言 |
| **CLI 框架** | [Cobra](https://github.com/spf13/cobra) | 命令行参数解析与子命令路由 |
| **区块链交互** | [go-ethereum](https://github.com/ethereum/go-ethereum) | 节点 RPC 通信、ABI 解析、日志 Topics 解码 |
| **关系型数据库** | MySQL | 持久化订单状态、地板价历史、索引进度（GORM） |
| **缓存/队列** | Redis | 状态缓存、分布式锁或撮合引擎的消息队列 |
| **日志** | [Uber Zap](https://github.com/uber-go/zap) | 结构化、高性能日志记录 |
| **配置管理** | Viper / YAML | 多环境配置加载 |

---

## 🏗 系统架构

系统主要分为三个核心层级：

1. **入口与调度层 (CLI & Config)**
    - 基于 Cobra 解析启动参数，加载多环境配置。
    - 初始化 Context、数据库连接池、Redis 客户端及日志实例。
2. **事件监听与解析层 (Listener & Parser)**
    - 轮询或订阅区块日志，根据合约 ABI 过滤目标事件（如 `LogCancel`）。
    - 将 `Topics` 中的 `indexed` 字段（如 `bytes32 OrderKey`）和 `Data` 中的非索引字段解码为 Go 结构体。
3. **持久化与维护层 (Persistence & Upkeep)**
    - **写入管道**：将解析后的事件组装为批量 SQL 参数,执行高效的批量 Upsert。
    - **Upkeep 循环**：独立协程运行定时任务（如每日清理过期地板价记录、定期拉取最新集合地板价），通过 `select` 监听 Ticker 与 `ctx.Done()` 实现无阻塞、可中断的后台维护。

---

## 🚀 快速开始

### 1. 环境要求
- Go 1.20 或更高版本
- MySQL 8.0+
- Redis 6.0+
- 可用的以太坊兼容节点 RPC 地址

### 2. 编译项目
```bash
# 克隆项目
git clone <your-repo-url>
cd MiSwapSync

# 编译二进制文件
go build -o miswapsync main.go