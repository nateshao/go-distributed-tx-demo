# SAGA（长事务补偿）Demo

**关键特性**：每步可补偿，可拆解

**适用场景**：复杂流程业务（如酒店+航班）

## 运行方法
```bash
cd saga
go run main.go
```

## 简介
本示例模拟 SAGA 分布式事务的正向和补偿操作，适合长流程、可拆解的业务场景。 