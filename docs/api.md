
Server's API (WIP)

## Basic concepts

The server convert tox callback to json format and forward to transport.
The server receive clients's RPC request of json format, execute it on tox instance and response clients.

The server support 2 transport type, Grpc and Websocket. 
They are standalone transport without depend on another one.

The server do more work other than just forward messages/events. It should make client simple.

## Transports

### Grpc

For RPC call, the result will response synchronizely.

For events push, server use Grpc's stream reponse model to do this work.

For this transport, client only need 1 connection.

### Websocket

For this transport, client need 2 connections.

One connection for RPC call. One connectoin for push events.

Client should not send any data to server over push connection.

Client should do exactly Request/Response pair on RPC connection. Client should handle MethodName*Resp* on this connection.

The 2 connections has different path.

### Comunicate flow

1. Client connect to RPC channel.
2. Call GetBaseInfo, get recently self info, friend list and group list.
3. Then client pull latest 20 history messages that just before GetBaseInfo call.
4. Client connect to pusher channel for wait later events.
5. When client running, send RPC calls over RPC api.

## Websocket API

### RPC requests

* connect path: /toxhsrpc

request format:
* Name: MethodName
* Args: array

response format:
* Name: MethodName
* Args: array

### pusher 

* connection path: /toxhspush

event format:
* Name: MethodName
* Args: array

All tox's callback will give a camel style of their original name, such as:

* tox\_callback\_friend\_message => FriendMessage
* tox\_callback\_conference\_message => ConferenceMessage

### GetBaseInfo
Call after connected to RPC, it's the first call to server.

request name: GetBaseInfo
arguments: ()

response format: https://github.com/envsh/tox-homeserver/blob/master/server/ths.proto#L24


### client pull history messages

request name: PullEventsByContactId
arguments: (prev_batch int)

response: array of Message struct https://github.com/envsh/tox-homeserver/blob/master/store/struct.go#L20

prev_batch: server will select early than prev\_batch. SQL: select ... where eventid <= prev\_batch order by eventid desc limit 20;

The first prev\_batch is returned by GetBaseInfo call.

this is for the features: loop get older messages

### Profile settings

coming soon...


