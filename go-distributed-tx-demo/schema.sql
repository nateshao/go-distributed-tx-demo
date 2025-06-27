-- 账户表
CREATE TABLE IF NOT EXISTS account (
  id INT PRIMARY KEY,
  balance INT NOT NULL
);

-- 库存表
CREATE TABLE IF NOT EXISTS inventory (
  id INT PRIMARY KEY,
  stock INT NOT NULL,
  reserved INT NOT NULL DEFAULT 0
);

-- 订单表
CREATE TABLE IF NOT EXISTS order_tbl (
  id INT PRIMARY KEY,
  status VARCHAR(20) NOT NULL
);

-- 消息表
CREATE TABLE IF NOT EXISTS message_tbl (
  order_id INT,
  status VARCHAR(20) NOT NULL,
  PRIMARY KEY(order_id)
); 