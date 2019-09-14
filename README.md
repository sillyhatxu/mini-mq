# mini-mq

## sample

[consumer](https://github.com/sillyhatxu/mini-mq/blob/master/client/consumer/consumer_test.go)

[producer](https://github.com/sillyhatxu/mini-mq/blob/master/client/producer/producer_test.go)


## develop backups

### Initialize your project

```
dep init
```

### Adding a dependency

```
dep ensure -add github.com/foo/bar github.com/foo/baz...

dep ensure -add github.com/foo/bar@1.0.0 github.com/foo/baz@master
```

### Updating dependencies

```
dep ensure -update github.com/sillyhatxu/logrus-client

dep ensure -update
```

# Release Template

#### Feature

* [NEW] Support for Go Modules [#17](https://github.com/sillyhatxu/convenient-utils/issues/17)

---

#### Bug fix

* [FIX] Truncate Latency precision in long running request [#17](https://github.com/sillyhatxu/convenient-utils/issues/17)

###

```
git tag v1.0.2
git push origin v1.0.2
```


### Initialize your project

```
go mod init github.com/sillyhatxu/mini-mq
```

### Updating dependencies

```
go mod vendor
```

### Upgrading dependencies
```
go get github.com/sillyhatxu/logrus-client@v2.0.3
```

### verify dependencies

```
go mod verify
```

### remove dependencies that is not used

```
go mod tidy
```

### print dependence diagram

```
go mod graph
```

### download dependencies

```
go mod download
```

# Release Template

#### Feature

* [NEW] Support for Go Modules [#17](https://github.com/sillyhatxu/convenient-utils/issues/17)

---

#### Bug fix

* [FIX] Truncate Latency precision in long running request [#17](https://github.com/sillyhatxu/convenient-utils/issues/17)


### 发布订阅(Publish/Subscribe)

> 使用sqlite做数据库
>
> 创建基础配置数据库和queue数据库
> 
> grpc做数据交互
>
> config设置数据每次consume数量
> 
> 批量提交（如何做到多个consume commit）
> 
> 设置commit超时时间
> 
> 
> 
> 

> https://github.com/grpc/grpc-go