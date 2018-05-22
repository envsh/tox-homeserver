### 消息同步

* 自己随便造的方式
* matrix协议的方式 https://matrix.org/docs/spec/client_server/r0.3.0.html#syncing
* ActiveSync方式

参考：z-push, matrix, 微信，


### 关于订阅的实现

* grpc+nats (现在的方式)
* grpc+custom
* comet
* websocket*2 (1 for rpc, 1 for push)
* grpc+emitter.io
* nrpc=rpc based on nats

### 参考：
pubsub库：https://github.com/lileio/pubsub
goim库

### FAQ

Q: android 防屏幕关闭JNI或者go实现

