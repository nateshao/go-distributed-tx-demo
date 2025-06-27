```mermaid
flowchart TD
  A["Client 请求"] --> B["Try 阶段：资源预留"]
  B --> C["Confirm 阶段：正式提交"]
  B --> D["Cancel 阶段：释放资源"]
``` 