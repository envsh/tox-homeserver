
 tox-homeserver

### Proposal

The main goal is make tox messages multiple endpoints synchronized and mobile friendly for save power and save network traffic. Then keep tox always online.

* Based on current toxcore protocol. Split out a server instance for personal use.
* It's support only one toxid. It's not a bouncer, but more like as a bridge.
* It's use gRPC/websocket for both support native app clients and web apps.
* The bridge store all messages when it's online. And all bridge's clients can receive messages synchronized .
* The clients can pull all history messages, and multiple clients can sync messages each other with bridge's help of cause.
* This is work in progress, still a demo, anyone interesting can take a look:
https://github.com/envsh/tox-homeserver
* The problem is now the group is temporary, there is not good way to merge group messages after recreate group.


### Run a server

https://github.com/envsh/tox-homeserver/releases

Download latest toxhs.tar.gz (for linux amd64)

    tar xvf toxhs.tar.gz   # file name is toxhs
    ./toxhs
    
Server will open 2 port: *:2080(Grpc), *:8099(Websocket+HTTP)

And you can see log line like this:
    
    2018-05-27 11:10:35.056859 I | [gofiat] xtox.go:68: ID: B5E7631D4D6C0EC0581B9DA34F431F0F2FAB8D115F9930D3DDA1333C8006A477A6440C224DD8

This is the server's(your's) tox id.

Warning: If you want recompile server, just use envsh/go-toxcore-c fork. There are 2 important fixes not PR to TokTok/go-toxcore-c.

### Run a native client

Goto release page, downoad a client for your OS:

    android arm: qofia-ffi-arm.apk
    windowns x64: qofia-ffi-amd64.tar.gz
    windows x32: qofia-ffi-i386.tar.gz
    linux x64: qofia-x86_64.AppImage
    linux x64(tared): qofia-linux-amd64-0.4.0.tar.bz2
    
For windows and linux clients, need seperate Qt installation.

Or download runtime simplely from https://github.com/qtchian/, and put .exe into bin directory:

* https://github.com/qtchina/qtenv_win64
* https://github.com/qtchina/qtenv_win32
* https://github.com/qtchina/qt510_qt_inline


### Web client (demo)

Open browser and visit: http://yourip:8099/webdui

### client snapshots

https://github.com/envsh/tox-homeserver/tree/master/snapshots

