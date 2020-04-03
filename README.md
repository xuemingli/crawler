# crawler
>>这是一个后端基于go语言的分布式爬虫项目，用来爬取某相亲网站里面的人物信息，并存储到Elasticsearch中，通过简单的前端页面进行筛选并显示。

## 环境
* Linux:
```Bash
    lxm@lxm:~$ cat /proc/version
    Linux version 5.3.0-42-generic (buildd@lcy01-amd64-019) (gcc version 7.4.0 (Ubuntu 7.4.0-1ubuntu1~18.04.1)) #34~18.04.1-Ubuntu SMP Fri Feb 28 13:42:26 UTC 2020
```
* golang:
```Bash
    lxm@lxm:~$ go version
    go version go1.13.5 linux/amd64
```
* docker:
```Bash
lxm@lxm:~$ sudo docker version
Client: Docker Engine - Community
 Version:           19.03.8
 API version:       1.40
 Go version:        go1.12.17
 Git commit:        afacb8b7f0
 Built:             Wed Mar 11 01:25:46 2020
 OS/Arch:           linux/amd64
 Experimental:      false

Server: Docker Engine - Community
 Engine:
  Version:          19.03.8
  API version:      1.40 (minimum version 1.12)
  Go version:       go1.12.17
  Git commit:       afacb8b7f0
  Built:            Wed Mar 11 01:24:19 2020
  OS/Arch:          linux/amd64
  Experimental:     false
 containerd:
  Version:          1.2.13
  GitCommit:        7ad184331fa3e55e52b890ea95e65ba581ae3429
 runc:
  Version:          1.0.0-rc10
  GitCommit:        dc9208a3303feef5b3839f4323d9beb36df0a9dd
 docker-init:
  Version:          0.18.0
  GitCommit:        fec368
```
* ElasticSearch:
```Bash
lxm@lxm:~$ sudo docker images
REPOSITORY                                      TAG                 IMAGE ID            CREATED             SIZE
docker.elastic.co/elasticsearch/elasticsearch   7.4.2               b1179d41a7b4        5 months ago        855MB
```

## 安装Docker和ElasticSearch
* 下载安装并启动docker
  * 免登录下载地址(Win)：https://download.docker.com/win/stable/Docker%20for%20Windows%20Installer.exe
  * 免登录下载地址(Mac)：https://download.docker.com/mac/stable/Docker.dmg
  * Linux上安装Docker(以Ubuntu为例):
```Bash
>>由于apt官方库里的docker版本可能比较旧，所以先卸载可能存在的旧版本：
$ sudo apt-get remove docker docker-engine docker-ce docker.io

>>更新apt包索引：
$ sudo apt-get update

>>安装以下包以使apt可以通过HTTPS使用存储库（repository）：
$ sudo apt-get install -y apt-transport-https ca-certificates curl software-properties-common

>>添加Docker官方的GPG密钥：
$ curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -

>>添加稳定版的源，使用下面的命令来设置stable存储库：
$ sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"

>>再更新一下apt包索引：
$ sudo apt-get update

>>列出可用的Docker-ce版本：
$ apt-cache madison docker-ce
 docker-ce | 5:19.03.5~3-0~ubuntu-bionic | https://download.docker.com/linux/ubuntu bionic/stable amd64 Packages
 docker-ce | 5:19.03.4~3-0~ubuntu-bionic | https://download.docker.com/linux/ubuntu bionic/stable amd64 Packages
 docker-ce | 5:19.03.3~3-0~ubuntu-bionic | https://download.docker.com/linux/ubuntu bionic/stable amd64 Packages
 docker-ce | 5:19.03.2~3-0~ubuntu-bionic | https://download.docker.com/linux/ubuntu bionic/stable amd64 Packages
 docker-ce | 5:19.03.1~3-0~ubuntu-bionic | https://download.docker.com/linux/ubuntu bionic/stable amd64 Packages
 docker-ce | 5:19.03.0~3-0~ubuntu-bionic | https://download.docker.com/linux/ubuntu bionic/stable amd64 Packages
 docker-ce | 5:18.09.9~3-0~ubuntu-bionic | https://download.docker.com/linux/ubuntu bionic/stable amd64 Packages
 docker-ce | 5:18.09.8~3-0~ubuntu-bionic | https://download.docker.com/linux/ubuntu bionic/stable amd64 Packages
 docker-ce | 5:18.09.7~3-0~ubuntu-bionic | https://download.docker.com/linux/ubuntu bionic/stable amd64 Packages
 docker-ce | 5:18.09.6~3-0~ubuntu-bionic | https://download.docker.com/linux/ubuntu bionic/stable amd64 Packages
 docker-ce | 5:18.09.5~3-0~ubuntu-bionic | https://download.docker.com/linux/ubuntu bionic/stable amd64 Packages
 docker-ce | 5:18.09.4~3-0~ubuntu-bionic | https://download.docker.com/linux/ubuntu bionic/stable amd64 Packages
 docker-ce | 5:18.09.3~3-0~ubuntu-bionic | https://download.docker.com/linux/ubuntu bionic/stable amd64 Packages
 docker-ce | 5:18.09.2~3-0~ubuntu-bionic | https://download.docker.com/linux/ubuntu bionic/stable amd64 Packages
 docker-ce | 5:18.09.1~3-0~ubuntu-bionic | https://download.docker.com/linux/ubuntu bionic/stable amd64 Packages
 docker-ce | 5:18.09.0~3-0~ubuntu-bionic | https://download.docker.com/linux/ubuntu bionic/stable amd64 Packages
 docker-ce | 18.06.3~ce~3-0~ubuntu | https://download.docker.com/linux/ubuntu bionic/stable amd64 Packages
 docker-ce | 18.06.2~ce~3-0~ubuntu | https://download.docker.com/linux/ubuntu bionic/stable amd64 Packages
 docker-ce | 18.06.1~ce~3-0~ubuntu | https://download.docker.com/linux/ubuntu bionic/stable amd64 Packages
 docker-ce | 18.06.0~ce~3-0~ubuntu | https://download.docker.com/linux/ubuntu bionic/stable amd64 Packages
 docker-ce | 18.03.1~ce~3-0~ubuntu | https://download.docker.com/linux/ubuntu bionic/stable amd64 Packages

>>安装最新版本的Docker CE：
$ sudo apt-get install -y docker-ce

>>选择要安装的特定版本，第二列是版本字符串，第三列是存储库名称，它指示包来自哪个存储库，以及扩展它的稳定性级别。要安装一个特定的版本，将版本字符串附加到包名中，并通过等号(=)分隔它们：
$ sudo apt-get install docker-ce=<VERSION>
例如：sudo apt-get install docker-ce=5:18.09.6~3-0~ubuntu-bionic

>>查看docker服务是否启动：
$ systemctl status docker

>>若未启动，则启动docker服务：
$ sudo systemctl start docker

>>测试经典的hello world：
$ sudo docker pull hello-world
$ sudo docker run hello-world
```
* 在docker中安装并运行elasticsearch镜像
  * 安装
  ```
  $docker pull docker.elastic.co/elasticsearch/elasticsearch:7.4.2
  ```
  * 运行
  ```
  $docker run -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" docker.elastic.co/elasticsearch/elasticsearch:7.4.2
  ```
## 安装go依赖包
  ```Bash
  go get github.com/olivere/elastic/v7
  ```
## 运行
### 单机版
* 运行go run crawler\main.go进行数据爬取
* 运行go run crawler\frontend\start.go启动本地web服务
* 访问localhost:8888
  * 首页信息如下图
  ![](https://github.com/xuemingli/crawler/blob/master/index.png "首页信息") 
  * 查询信息如下图
  ![](https://github.com/xuemingli/crawler/blob/master/show.png "信息展示") 
### 分布式版
* 启动ElasticSearch存储服务，RPC服务于本地的1234端口：
  ```Bash
  go run crawler\crawler_distributed\persist\server\itemSaver.go --port=1234
  ```
* 启动多个worker,用不同的端口进行RPC服务：
  ```Bash
  go run crawler\crawler_distributed\worker\server\worker.go --port=9000
  go run crawler\crawler_distributed\worker\server\worker.go --port=9001
  go run crawler\crawler_distributed\worker\server\worker.go --port=9002
  ```
* 启动engine，进行主程序调度，包括一个存储服务的RPC客户端协程：
  ```Bash
  go run crawler\crawler_distributed\main.go --itemsaver_host=":1234" --worker_host=":9000,:9001,:9002"
  ```
