```mermaid
flowchart TD
  A["Saga Step1"] --> B["Saga Step2"]
  B -- 失败 --> C["补偿 Step2"]
  C --> D["补偿 Step1"]
  B -- 成功 --> E["流程结束"]
``` 