```mermaid
flowchart TD
  A["分布式事务开始"] --> B["分库分表执行"]
  B --> C["两阶段/三阶段协议"]
  C --> D["所有节点提交"]
  C --> E["部分节点失败，回滚"]
``` 