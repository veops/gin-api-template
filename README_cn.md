<h3 align="center">基于gin框架的api项目模板</h3>
<div align="center">

[![Licence](https://img.shields.io/badge/Licence-Apache-brightgreen)](https://github.com/veops/gin-api-template/blob/main/LICENSE)
[![API](https://img.shields.io/badge/API-gin-blue)](https://github.com/gin-gonic/gin)
[![Log](https://img.shields.io/badge/Log-zap-green)](https://github.com/uber-go/zap)
[![Golang](https://img.shields.io/badge/go-1.18+-blue)](https://go.dev/dl/)
</div>

------------------------------

[English](README.md)
## 概览

这是一个基于gin框架的简单项目建设模板。如果你想快速使用golang构建后端项目，这个模板非常适合你。
你可以直接使用该项目，并在此基础上快速开发你自己的项目。

## 项目架构

``` shell
├── cmd
│   ├── apps # 子项目目录`
│   │   ├── config.example.yaml # 启动的配置文件样例
│   │   └── server.go # 某子项目启动
│   └── main.go # 项目启动的入口
├── docs # 各种文档
└── pkg  # 项目核心代码包 
    ├── conf
    │   └── conf.go # 全局的配置设置 
    ├── docs # 各种文档
    ├── logger # 日志的设置
    ├── server # 项目的核心逻辑块， 如果是多个子模块，可以建多个不同的目录，这里示例中建立一个server目录
    │   ├── auth # 认证模块 
    │   │   └── acl # 默认的acl认证
    │   ├── controller # 控制器模块
    │   │   ├── controller.go # 一个全局controller的定义
    │   │   └── hello.go # 一个api样例，每种类型的api接口，建立一个单独的文件，这样条理比较清晰 
    │   ├── model # 存储结构的配置，各种数据库存储的model结构配置，如数据库的各个字段的定义等
    │   │   └── hello.go
    │   ├── router # 路由的各种定义
    │   │   ├── middleware # 各种中间件的的定义
    │   │   │   ├── auth.go # 认证
    │   │   │   ├── cache.go # 缓存
    │   │   │   ├── log.go # 日志
    │   │   │   └── ... 
    │   │   ├── router.go # 路由全局配置
    │   │   └── routers.go # 各种路由的配置,api代码逻辑主要在这里配置
    │   └── storage # 后端存储的实现
    │       ├── cache # 缓存实现
    │       │   ├── local # 内存
    │       │   └── redis # reids
    │       └── db # 各种数据库的存储
    │           └── mysql
    └── util # 放置各种公用的函数、工具等
        └── util.go
```

## 特点
- 项目启动使用 **cobra** 进行命令行启动。
- 使用 **yaml** 格式项目的配置文件
- 日志采用zap进行日志输出，封装了 **日志输出**，**日志轮转**, **动态修改日志级别** 等
- 封装 **路由配置** router
- 封装了多个 **middleware**

## 快速开始
### 第一步  克隆项目
```sh
git clone git@github.com:veops/gin-api-template.git
```
### 第二步  修改配置
```sh
cd gin-api-template/cmd
cp apps/config.example.yaml apps/config.yaml 
# 修改config.yaml中的配置
```

### 第三步 构建、运行项目
```
go build -o server main.go 
./server run -c apps
```

### 第四步 验证
> 以下验证任意一个正常
- **服务可用性检查**, **作为API服务必须具备**
```
curl -X GET http://localhost:8080/-/health
```
返回: `OK`
- **动态修改日志级别**。这个接口方便在遇到问题时，无需重启服务，动态调整日志的输出级别，方便排查问题。 
```
curl -X PUT localhost:8080/-/log/level -H "Content-Type: application/json" -d '{"level":"debug"}'
```
返回:
```json
{"level":"debug"}
```
- 普通api验证
```shell
curl -X GET http://localhost:8080/api/v1/hello
```
 返回:
```json
{
  "code":0,
  "data":{
    "Time":"2023-11-06T09:36:55.830076+08:00"
  },
  "message":"hello world"
}
```

## 常见问题解答
### 如何添加新的路由（例如API）？
> 1. 在pkg/server/controller中编写你自己的处理程序，就像hello.go所做的那样。
> 2. 将你的处理程序注册到pkg/server/router/routes.go中的路由中。 
### 如何添加新的中间件？
> 1. 在pkg/server/middleware中默认提供了一些中间件。如果你没有找到需要的中间件，可以在这里编写自己的中间件。
> 2. 然后你应该在pkg/server/router的setupRouter()函数中注册你自己的中间件。
### 如何使用自己的授权方式？
> 1. 在pkg/server/router/middleware/auth.go中添加你自己的授权方式。默认提供了ACL、白名单和基本身份验证。
> 2. 如果需要的话，在Auth()函数中重新排列授权函数的顺序
### 什么是ACL？
> ACL 是一个基于角色的资源权限管理服务,详细可参考 [https://github.com/veops/acl](https://github.com/veops/acl)
### 如何使用自己喜欢的数据库？
> 1. 由于这个项目是一个通用模板，不同情况下数据库的选择可能会有很大差异，因此我们没有提供默认的数据库。mysql文件夹仅用于演示。
> 2. 假设你想添加Mongo数据库。你可以在db文件夹中添加一个名为mongo.go的文件，并完成Mongo的初始化逻辑。
> 3. 在conf/conf.go中添加Mongo配置结构体。
> 4. 在你的config.yaml文件中添加Mongo配置。




---
_**欢迎关注公众号(维易科技OneOps)，关注后可加入微信群，进行产品和技术交流。**_

![公众号: 维易科技OneOps](docs/images/wechat.jpg)

