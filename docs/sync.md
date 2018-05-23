### 消息同步

* 自己随便造的方式
* matrix协议的方式 https://matrix.org/docs/spec/client_server/r0.3.0.html#syncing
* ActiveSync方式

参考：z-push, matrix, 微信，

### 消息同步流程

同步并存储流程

1. 启动app，拉取每个房间的最新n条消息。
2. 不断滚动加载更旧消息。
3. 实时消息不断更新拉取点记录。

房间：可以是好友或群

### 关于时间线

有3条时间线：
1. timeline1: server永久存储端时间线
2. timeline2: client 界面时间线
3. timeline3: client永久存储时间线

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
A: done with heavy wrapper, not just binding.

