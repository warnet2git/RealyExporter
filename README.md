# 中继导出器 RelayExporter

## 背景
开发`Prometheus`格式的数据生成功能时，在本地调式不是一件很方便的事情。通常可行的办法有：
1. 增加`/metrics`接口，将数据在这个接口暴露给`Prometheus`；
2. 通过`Pushgateway`将数据推给`Prometheus`；

从个人的实践来看：
1. 方法1存在的问题是：需要较多的修改业务代码，侵入性比较强；
2. 方法2存在的问题是：不支持时间戳，不方便核对数据，验证正确性；

所以，做了个小工具，方便将数据导入到`Prometheus`来观察数据的生成情况。

## 使用方法
启动服务的方法`go run RelayExporter.go`
注意：
1. 服务启动的端口配置在参数`ServerAddr`上，默认9123端口；
2. 将需要查看的数据`POST`到`/add`接口；
3. 将`Prometheus`的scrape配置到`ServerAddr`对应端口的`/metrics`接口上；
4. 下面给出简单的验证方法
```shell
    curl -X POST -d "MyMetric1{key1=\"val1\",key2=\"val2\"} 123 1625886000000" "http://localhost:9123/add"
    curl "http://localhost:9123/metrics"
```

## 其他问题
1. `Prometheus`在抓取数据的时候，可能出现错误，为了方便定位错误，可以增加启动参数`--log.level=debug`，这样就可以看到错误信息，但是错误信息中不会展示时间戳；
2. `Prometheus`项目的地址：https://prometheus.io/
3. `Grafana`项目的地址：https://grafana.com/

## 最后
enjoy :)
