-- lightsocks服务器列表
CREATE TABLE server (
  ip       VARCHAR(39) PRIMARY KEY, --服务器外网IP
  port     INT       NOT NULL, --服务器RPC监听端口
  alive    BOOL               DEFAULT FALSE, --是否可以访问当前服务器的服务
  status   JSONB, --服务器状态
  updateat TIMESTAMP NOT NULL DEFAULT now() --最后更新的时间,
);

-- lightsocks服务列表
CREATE TABLE service (
  addr     VARCHAR(45) PRIMARY KEY, --id由ip和port拼接成 ip:port
  password VARCHAR(344) NOT NULL, --lightsocks服务密码
  ip       VARCHAR(39) REFERENCES server ON DELETE CASCADE--服务器外网IP
);

-- 注册的用户
CREATE TABLE "user" (
  email    VARCHAR(30) PRIMARY KEY, --邮件
  password BYTEA               NOT NULL, --加密后的密码
  token    VARCHAR(172) UNIQUE NOT NULL, --鉴权token
  balance  FLOAT DEFAULT 0, --用户余额 单位分
  cost     FLOAT DEFAULT 0, --用户每日花费 单位分
  config   JSONB --用户配置
);

-- 用户使用服务关系表
CREATE TABLE user_service (
  email  VARCHAR(30) REFERENCES "user" ON DELETE CASCADE, --user邮件
  addr   VARCHAR(45) REFERENCES service ON DELETE CASCADE, --service id
  uuid   VARCHAR(36) UNIQUE, --唯一的机器UUID
  device JSONB--机器的信息
);