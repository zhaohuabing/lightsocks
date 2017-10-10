[![Build Status](https://img.shields.io/travis/gwuhaolin/lightsocks.svg?style=flat-square)](https://travis-ci.org/gwuhaolin/lightsocks)

# Lightsocks
一个轻量级网络混淆代理，基于 SOCKS5 协议，可用来代替 Shadowsocks。
采用更高效的算法，专注于翻墙。

## 安装
去 [releases](https://github.com/gwuhaolin/lightsocks/releases) 页下载最新的可执行文件，注意选择正确的操作系统和位数。
解压后会看到2个可执行文件，分别是：

- **lightsocks-local**：用于运行在本地电脑的客户端，用于桥接本地浏览器和远程代理服务，传输前会混淆数据；
- **lightsocks-server**：用于运行在墙外服务器的客户端，会还原混淆数据；

## 启动
在墙外服务器下载好 lightsocks-server 后，执行命令：
```bash
./lightsocks-server [./path/to/config.json]
```
就可启动服务端。


在本地电脑下载好 lightsocks-local 后，执行命令：
```bash
./lightsocks-local [./path/to/config.json]
```
就可启动本地代理客户端。

## 配置
lightsocks-local 支持的选项：
- **password**：用于加密数据的密码，字符串格式，在没有填时会自动生成；
- **listen**：本地 SOCKS5 代理客户端的监听地址，格式为 `ip:port`，默认为 `0.0.0.0:7474`；
- **remote**：墙外服务器的监听地址，格式为 `ip:port`，默认为 `0.0.0.0:7474`。

lightsocks-server 支持的选项：
- **password**：用于加密数据的密码，字符串格式，在没有填时会自动生成；
- **listen**：本地 SOCKS5 代理客户端的监听地址，格式为 `ip:port`，默认为 `0.0.0.0:7474`。

注意：
- lightsocks-local 和 lightsocks-server 的 password 必须一致才能正常翻墙，password 不要泄露。
- password 会自动生成，不要自己生成。 


只能通过 JSON 文件的方式传参给 lightsocks-local 和 lightsocks-server，
例如启动时执行命令 `./lightsocks-local config.json`，则 `config.json` 文件内容需如下：
```json
{
  "remote": "45.56.76.5:7474",
  "password": "..."
}
```





