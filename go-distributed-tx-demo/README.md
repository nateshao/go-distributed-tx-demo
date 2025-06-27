# go-distributed-tx-demo

本项目包含六种主流分布式事务方案的最小可运行 Go Demo，适合学习和对比分布式事务的实现原理与适用场景。

## 目录结构
- twopc/        —— 2PC（两阶段提交）
- tcc/          —— TCC（Try-Confirm-Cancel）
- msgtable/     —— 本地消息表/事务消息
- saga/         —— SAGA（长事务补偿）
- reliablemsg/  —— 可靠消息+状态检查
- strongdb/     —— 强一致数据库分布式事务

## 各方案简介与对比
| 类型 | 关键特性 | 适用场景 | 主要优缺点 |
|------|----------|----------|------------|
| 2PC | 原子性好，阻塞，性能差 | 银行、支付等强一致性 | 强一致但易阻塞，单点故障风险 |
| TCC | 业务级接口设计，高性能 | 电商下单/库存/优惠券 | 需业务侵入，接口复杂 |
| 本地消息表 | 最终一致性，常见于消息队列 | 微服务解耦、异步下单 | 实现简单，延迟一致性 |
| SAGA | 每步可补偿，可拆解 | 复杂流程业务 | 需补偿逻辑，弱一致性 |
| 可靠消息 | 配合 MQ 事务，灵活 | RocketMQ、Kafka | 依赖 MQ，需状态回查 |
| 强一致数据库 | 原生分布式事务 | TiDB、CockroachDB | 依赖数据库，性能受限 |

## 运行方法
进入任一子目录，执行：
```bash
cd <子目录>
go run main.go
```

## 流程图
每个子目录下均有流程图（Mermaid 格式），可用 VSCode 插件或在线工具查看。

## 参考资料
- [分布式事务原理与实践](https://zhuanlan.zhihu.com/p/34823738)
- [阿里云分布式事务解决方案](https://help.aliyun.com/document_detail/155460.html)
- [TiDB 分布式事务](https://docs.pingcap.com/zh/tidb/stable/distributed-transaction)
- [SAGA Pattern](https://microservices.io/patterns/data/saga.html)
- [2PC & TCC 论文](https://cs.brown.edu/people/pvh/Talks/2pc.pdf) 