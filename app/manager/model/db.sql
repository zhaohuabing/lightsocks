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
  ip       VARCHAR(49) REFERENCES SERVER ON DELETE CASCADE, --服务器外网IP,
  port     INT   NOT NULL, --lightsocks服务监听端口
  password BYTEA NOT NULL, --lightsocks服务密码
  UNIQUE (ip, port)
);