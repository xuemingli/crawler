# crawler
>>这是一个用go语言写的爬虫项目，用来爬取 http://www.zhenai.com 网站里面的人物信息，将信息存储到Elasticsearch中，通过简单的前端页面进行筛选并显示。

## 环境
* golang: `v1.13`
* docker: `19.03.5`
* ElasticSearch: `7.0`

## 运行
* 下载安装并启动docker
  * 免登录下载地址(Win)：https://download.docker.com/win/stable/Docker%20for%20Windows%20Installer.exe
  * 免登录下载地址(Mac)：https://download.docker.com/mac/stable/Docker.dmg
  * 在命令行输入`docker version`查看版本信息，我的如下所示：
  ```Bash
  Client: Docker Engine - Community
   Version:           19.03.5
   API version:       1.40
   Go version:        go1.12.12
   Git commit:        633a0ea
   Built:             Wed Nov 13 07:22:37 2019
   OS/Arch:           windows/amd64
   Experimental:      false
 
  Server: Docker Engine - Community
   Engine:
    Version:          19.03.5
    API version:      1.40 (minimum version 1.12)
    Go version:       go1.12.12
    Git commit:       633a0ea
    Built:            Wed Nov 13 07:29:19 2019
    OS/Arch:          linux/amd64
    Experimental:     false
  containerd:
   Version:          v1.2.10
   GitCommit:        b34a5c8af56e510852c35414db4c1f4fa6172339
  runc:
   Version:          1.0.0-rc8+dev
   GitCommit:        3e425f80a8c931f88e6d94a8c831b9d5aa481657
  docker-init:
   Version:          0.18.0
   GitCommit:        fec3683
  ```
* 在docker中运行elasticsearch镜像
* 运行main.go进行数据爬取
* 运行start.go启动本地web服务
* 访问localhost:8888
