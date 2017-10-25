### 一、Build 镜像

lightsocks 的 docker 镜像采用当前目录下的 Dockerfile 创建，在安装有 docker 环境的机器上 build 即可

``` sh
git clone https://github.com/gwuhaolin/lightsocks.git
cd lightsocks/docker
docker build -t lightsocks .
```

### 二、Docker 镜像说明

该镜像同时集成了 server 端和 client 端，默认配置文件保存在 `/root/.lightsocks.json` 中

server 端与 client 端切换可以通过更改 `LIGHT_MODULE` 环境变量值实现，该变量默认值为 `lightsocks-server` 即服务端，
将其修改为 `lightsocks-local` 即以客户端方式运行，同时镜像也支持直接运行命令的方式切换，具体使用如下

#### 2.1、服务端

**由于 lightsocks 的配置文件是自动生成的，所以需要先让服务端帮我们生成一个配置文件，出现密码后 Ctrl+c 停止即可**

``` sh
docker run --rm -it -v `pwd`:/root mritd/lightsocks:1.0.3
```

**查看 server 端生成的 `.lightsocks.json` 中的端口和密码**

``` sh
cat .lightsocks.json

{
    "listen": ":43413",
    "remote": "",
    "password": "vfpfHaoIliU7odM036nYvLXolOn8hb53dtfFu9vtZ530le9ywWsBfHomaLmCiyhvQRsKPVWn1A9DW/OOpGWuLAyQAwvdFmQ2tvkVik+Y8SqSLQUN2TD4E3irOTOJTqBU3OrOISRRAvVN/WBWI9F+FMKcccNjIoQY7kddEfCPppNSysdIrH9tPkJckXk4EFBz3uZEuBqjgSscRln+bEW6PzxqWIduSa2yIA7AMajnl17ymvd7Euxm0st0cLcvJy4fiDfPUwfkg7SGBtCMacml69oy4p/jm6/ESwDWgKJXfb9iNUDN5QTM4cbVGY11SrEJ+7Ow9h6eYVqZF0wp4P/IOg=="
}
```

**重新启动 server 端并增加 docker 的端口映射**

``` sh
docker run -d --name lsserver -v `pwd`:/root -p 43413:43413 --restart=always mritd/lightsocks:1.0.3 lightsocks-server
```

#### 2.2、客户端

从服务端的配置中已经可以拿到端口和密码，客户端直接编写配置启动即可

``` sh
cat > .lightsocks.json <<EOF
{
    "listen": ":43413",
    "remote": "SERVER_IP:43413",
    "password": "vfpfHaoIliU7odM036nYvLXolOn8hb53dtfFu9vtZ530le9ywWsBfHomaLmCiyhvQRsKPVWn1A9DW/OOpGWuLAyQAwvdFmQ2tvkVik+Y8SqSLQUN2TD4E3irOTOJTqBU3OrOISRRAvVN/WBWI9F+FMKcccNjIoQY7kddEfCPppNSysdIrH9tPkJckXk4EFBz3uZEuBqjgSscRln+bEW6PzxqWIduSa2yIA7AMajnl17ymvd7Euxm0st0cLcvJy4fiDfPUwfkg7SGBtCMacml69oy4p/jm6/ESwDWgKJXfb9iNUDN5QTM4cbVGY11SrEJ+7Ow9h6eYVqZF0wp4P/IOg=="
}
EOF
docker run -d --name lsclient -v `pwd`:/root -p 43413:43413 --restart=always mritd/lightsocks:1.0.3 lightsocks-local
```

### 三、使用 docker-compose 启动

纯命令行的方式启动可能并不友好，以下为一个服务端客户端通用的 docker-compose 文件

``` sh
version: '2' 
services:
  lightsocks:
    image: mritd/lightsocks:1.0.3
    restart: always
    volumes: 
      - ./:/root
    ports: 
      - "43413:43413"
    environment:
      - LIGHT_MODULE
```

**docker-compose 启动前需要先 `export LIGHT_MODULE=xxxx` 来指定运行模式；
同样 server 启动前需要先启动一次生成配置文件，然后再修改 docker-compose 的端口映射后启动，
关于 docker-compose 安装及使用这里不再详细阐述**
