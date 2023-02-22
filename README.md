# mini-tiktok-backend
项目源自字节跳动青训营（后端方向），是结营提交的最终项目成果。项目采用Golang编写，项目结构清晰，代码符合编码规范。项目实现了基础功能、互动方向与社交方向，在客户端下正常运行。在定义IDL时采用参数校验，边界情况处理良好。除此之外，进行了大量的性能优化，包括架构层面与语言层面的优化等。安全性方面，使用jwt验证授权，以及OSS签名链接等。
<a name="q3J0q"></a>
# 项目特点

- 分布式微服务架构
- 使用hertz（高性能HTTP框架）、kitex（高性能RPC框架）、gorm
- 使用thrift IDL定义接口，并用hz和kitex生成项目代码
- 使用etcd进行服务注册与发现
- 使用go-tagexpr和thrift-gen-validator插件进行参数校验
- 使用jwt进行用户验证授权
- 数据库使用MySQL，对象存储使用阿里云OSS，Redis负责数据缓存
- 使用opentelemetry和jaeger进行链路追踪及其可视化
- 使用grafana、jaeger、pprof进行数据监控
- 使用hlog和klog进行日志处理
- 使用docker compose负责项目中容器集群的快速编排
<a name="xPxik"></a>
# 项目架构
![architecture.png](https://cdn.nlark.com/yuque/0/2023/png/12760556/1676980652173-530989e0-7da7-4b49-a86c-e81ae72227a1.png#averageHue=%23f6f6f6&clientId=u13d39cde-b026-4&from=drop&id=u853db267&name=architecture.png&originHeight=1025&originWidth=1266&originalType=binary&ratio=1.5625&rotation=0&showTitle=false&size=1077792&status=done&style=none&taskId=u5cd658e4-ba42-4767-9d08-4615a810b42&title=)
<a name="qWRO3"></a>
# 调用关系
![屏幕截图 2023-02-20 225939.png](https://cdn.nlark.com/yuque/0/2023/png/12760556/1676980656565-d5b293f5-a375-447b-a758-c387505d9674.png#averageHue=%23f6f6f6&clientId=u13d39cde-b026-4&from=drop&id=uddb2f726&name=%E5%B1%8F%E5%B9%95%E6%88%AA%E5%9B%BE%202023-02-20%20225939.png&originHeight=720&originWidth=957&originalType=binary&ratio=1.5625&rotation=0&showTitle=false&size=118585&status=done&style=none&taskId=u9a92c10b-393e-47d6-a144-fef22d90d7e&title=)
<a name="o7zeq"></a>
# 项目结构
```c
mini-tiktok-backend/
├── app # 抖声客户端
├── cmd
│   ├── api # api服务
│   │   ├── biz
│   │   │   ├── handler/ # http请求处理
│   │   │   ├── model/ # hz生成代码
│   │   │   ├── mw/ # 中间件
│   │   │   ├── router/ # 路由与中间件配置
│   │   │   └── rpc/ # 发起的rpc调用
│   ├── comment # comment服务
│   │   ├── dal/ # 数据库操作
│   │   ├── pack/ # 包装返回数据
│   │   ├── rpc/ # 发起的rpc调用
│   │   ├── service/ # 具体服务逻辑
│       └── handler.go # 提供的服务
├── idl
│   ├── api/ # api相关idl，用于hz
│   └── service/ # service相关idl，用于kitex
├── kitex_gen/ # kitex生成代码
├── pkg
│   ├── configs/ # otel和sql配置
│   ├── consts/ # 常量集合
│   ├── errno/ # 错误码
│   └── mw/ # 公共中间件
├── test # 测试文件
├── docker-compose.yaml # docker compose模板文件
└── README.md
```
<a name="zUkaC"></a>
# 数据库![db.png](https://cdn.nlark.com/yuque/0/2023/png/12760556/1676980664780-c8e77ddc-65d5-4617-b109-52f4506a1622.png#averageHue=%23fbfafa&clientId=u13d39cde-b026-4&from=drop&id=u6b84181a&name=db.png&originHeight=656&originWidth=992&originalType=binary&ratio=1.5625&rotation=0&showTitle=false&size=239337&status=done&style=none&taskId=uc4720c45-736c-45c0-b63c-cb6a8937396&title=)

<a name="bheAi"></a>
# 链路追踪
![屏幕截图 2023-02-12 142258.png](https://cdn.nlark.com/yuque/0/2023/png/12760556/1676182994708-a438f3d0-09a4-4fda-9753-c78cfde347b4.png#averageHue=%23f8f4f1&clientId=uddb70f32-1692-4&from=drop&id=u15a29ff5&name=%E5%B1%8F%E5%B9%95%E6%88%AA%E5%9B%BE%202023-02-12%20142258.png&originHeight=792&originWidth=1898&originalType=binary&ratio=1.5625&rotation=0&showTitle=false&size=122013&status=done&style=none&taskId=u37624a37-006d-49a3-bb08-54f4b15c8ba&title=)
<a name="FaUq3"></a>
# 性能优化
采用的性能优化方法有：

- gzip压缩优化性能
- redis存储计数数据
- SQL预编译
- Gorm跳过默认事务
- rpc连接多路复用
- 数据库设置索引
- 字符串处理优化

实测在高并发场景下（受于机器限制，使用的并发量是100），优化性能后的请求处理的平均耗时缩短的40%。<br />![屏幕截图 2023-02-12 142019.png](https://cdn.nlark.com/yuque/0/2023/png/12760556/1676182834716-f953acfd-8ceb-4563-8fcb-d03927b87078.png#averageHue=%23fefefe&clientId=uddb70f32-1692-4&from=drop&id=uab5c79dd&name=%E5%B1%8F%E5%B9%95%E6%88%AA%E5%9B%BE%202023-02-12%20142019.png&originHeight=327&originWidth=1478&originalType=binary&ratio=1.5625&rotation=0&showTitle=false&size=33249&status=done&style=none&taskId=u5c73927f-97bf-4822-a517-b72e905a2b8&title=)![屏幕截图 2023-02-12 141143.png](https://cdn.nlark.com/yuque/0/2023/png/12760556/1676182834717-8d25acc1-15ad-415e-961e-366370a5b048.png#averageHue=%23fefefe&clientId=uddb70f32-1692-4&from=drop&id=uff7a3052&name=%E5%B1%8F%E5%B9%95%E6%88%AA%E5%9B%BE%202023-02-12%20141143.png&originHeight=322&originWidth=1477&originalType=binary&ratio=1.5625&rotation=0&showTitle=false&size=32876&status=done&style=none&taskId=u3564a00c-273c-4158-bf1b-824373fb206&title=)
<a name="VY9Xy"></a>
# 安全性

- jwt验证授权
- oss签名链接
- 密码加密存储
- SQL预编译
<a name="XLsJY"></a>
# 项目部署
```bash
// 启动容器集群
docker-compose up

// 启动api服务
cd cmd/api
go run .

// 启动comment服务
cd cmd/comment
sh build.sh
sh output/bootstrap.sh

// 启动favorite服务
cd cmd/favorite
sh build.sh
sh output/bootstrap.sh

// 启动publish服务
cd cmd/publish
sh build.sh
sh output/bootstrap.sh

// 启动user服务
cd cmd/user
sh build.sh
sh output/bootstrap.sh

// 启动video服务
cd cmd/video
sh build.sh
sh output/bootstrap.sh

// 启动relation服务
cd cmd/relation
sh build.sh
sh output/bootstrap.sh
```
