# xinsitute_platform


## 基于eagle框架（gin+gorm+viper+sarma）


### 生成pb文件
- 为避免不同开发同学的protoc版本不一致，统一使用make pb生成，依赖docker镜像
```shell
make pb
```

### 生成grom的model代码
- mysql —> go gorm 推荐用在线工具转换
- 数据库crud操作使用推荐dbhelper


### kafka
- https://github.com/asong2020/Golang_Dream/tree/master/code_demo/kafka_demo
- 基建：infrastructure/queue/kafka/

### 七牛云图片服务
- infrastructure/utils/url.go

### 日志服务
logger


### 服务划分
admin作为服务分发rpc服务层
core作为rpc服务核心
