# crawler
===
这是一个用go语言写的爬虫项目，用来爬取 http://www.zhenai.com 网站里面的人物信息，将信息存储到Elasticsearch中，通过简单的前端页面进行筛选并显示。

# 环境
---
golang: v1.13
docker: 19.03.5
ElasticSearch: 7.0

# 运行
---
1.启动docker
2.在docker中运行elasticsearch镜像
3.运行main.go进行数据爬取
4.运行start.go启动本地web服务
5.访问localhost:8888
