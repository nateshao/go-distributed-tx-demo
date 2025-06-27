# 可靠消息 + 状态检查 Demo

**关键特性**：配合 MQ 事务，灵活

**适用场景**：RocketMQ、Kafka 等异步事务

## 运行方法
```bash
cd reliablemsg
go run main.go
```

## 简介
本示例模拟可靠消息发送和状态回查的流程，适合与消息队列结合的分布式事务场景。 