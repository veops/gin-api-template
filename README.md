<h3 align="center">A simple project template based on gin framework</h3>
<div align="center">

[![Licence](https://img.shields.io/badge/Licence-Apache-brightgreen)](https://github.com/veops/gin-api-template/blob/main/LICENSE)
[![API](https://img.shields.io/badge/API-gin-blue)](https://github.com/gin-gonic/gin)
[![Log](https://img.shields.io/badge/Log-zap-green)](https://github.com/uber-go/zap)
[![Golang](https://img.shields.io/badge/go-1.18+-blue)](https://go.dev/dl/)
</div>

------------------------------

[中文](README_cn.md)
## Overview
A simple project construction template based on the gin framework.

If you want to quickly build a backend project using golang, this template is perfect for you.
You can directly use this project and quickly develop your own project based on it.

## Architecture Overview


``` shell
├── cmd
│   ├── apps # Subdirectory for sub-projects`
│   │   ├── config.example.yaml # Example configuration file for startup
│   │   └── server.go # Starts a specific sub-project`
│   └── main.go # The entpoint for starting the project.
└── pkg  # Core code package of the project`
    ├── conf
    │   └── conf.go # Global configuration settings` 
    ├── docs # various documents
    ├── logger # Logging settings
    ├── server # Core logic block of the project. If there are multiple sub-modules, multiple directories can be created. Here is an example of a server directory is provided.
    │   ├── auth # Authentication module` 
    │   │   └── acl # Default ACL authentication
    │   │   └── ... # Any other authentication can be placed in a separate directory, such as ldap 
    │   ├── controller # Controller module`
    │   │   ├── controller.go # Definition of a global controller
    │   │   └── hello.go # An example API. Each type of API interface should have a separate file for better organization. 
    │   ├── model # Configuration of storage structures, including model configurations for various database storages, such as definitions of database fields, etc.
    │   │   └── hello.go
    │   ├── router # Various router definitions
    │   │   ├── middleware # Definition of various middlewares
    │   │   │   ├── auth.go # Authentication
    │   │   │   ├── cache.go # Cache
    │   │   │   ├── log.go # Logging
    │   │   │   └── ... 
    │   │   ├── router.go # Global router configuration
    │   │   └── routers.go # Configuration of various routes. The main logic of the API code is configured here.
    │   └── storage # Implementation of backend storage`
    │       ├── cache # Cache implementation` 
    │       │   ├── local # In-memory storage
    │       │   └── redis # Redis storage
    │       └── db # Storage for various databases
    │           └── mysql
    └── util # Contains various common functions and utilities
        └── util.go
```

## Features
- Cobra for command-line startup.
- Uses YAML format for configuration files.
- Logging is done using Zap, which provides logging output and log rotation functionalities.
- Encapsulates router configuration.
- Encapsulates multiple middlewares.

## Getting Started
### Step 1. Clone the project
```sh
git clone git@github.com:veops/gin-api-template.git

```
### Step 2. Modify config file
```sh
cd gin-api-template/cmd
cp apps/config.example.yaml apps/config.yaml 
# modify config.yaml
```
### Step 3. Build and run the project
```
go build -o server main.go 
./server run -c apps
```
### Step 4. Validation
- Access address `http://localhost:8080/-/health` by Browser or terminal, the response is `OK`
- Access address `http://localhost:8080/api/v1/hello` the response as follows: 
```json
{
  "code":0,
  "data":{
    "Time":"2023-11-06T09:36:55.830076+08:00"
  },
  "message":"hello world"
}
```

## FAQ

### How can I add new routes(i.e. api)
> 1. Write your own handler in pkg/server/controller, just like what hello.go does
> 2. Register your handler to routes in pkg/server/router/routes.go

### How can I add new middlewares
> 1. There all some middlewares providered by default in pkg/server/middlware, therefore, plaease check them firstly. If you do not find what you need, just write your own one here.
> 2. Then you should register yours in the setupRouter() function in pkg/server/router

### How to use my own authorization
> 1. Add your own authorization way in pkg/server/router/middleware/auth.go. ACL, white list and basic auth is provided by default.
> 2. Rearrange the order of auth functions if you need in Auth() function

### What is ACL

> 1. A role-based resource permission management service. Please check defails here [https://github.com/veops/acl](https://github.com/veops/acl)

### How to use my prefered database
> 1. Since this project is a common template and cohices of db is much different in different conditions, we did not provide a default one, the folder of mysql is just used to demonstrate.
> 2. Suppose you want to add mongo database. You can add a folder in db with a file named mongo.go in it. Then you should finish init logic of mongo.
> 3. Add mongo config struct in conf/conf.go
> 4. Add mongo config in your config.yaml file