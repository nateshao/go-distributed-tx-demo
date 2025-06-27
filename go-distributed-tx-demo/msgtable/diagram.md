```mermaid
flowchart TD
  A["本地事务"] --> B["写入本地消息表"]
  B --> C["异步投递消息"]
  C --> D["消费方处理"]
``` 