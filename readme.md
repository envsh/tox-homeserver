# tox-homeserver

### Proposal

The main goal is,

1. To make tox messages synchronized between multiple endpoints,
2. Mobile friendly on saving power and saving network traffic,
3. Keep toxid always online.

The main purposes are,

* Based on toxcore protocol. Split out a server instance for personal use.
* Supports only one toxid. It's not a bouncer, but more act as a bridge.
* Supports native apps and web apps by gRPC/websocket.
* Store messages and synchronize messages to bridged clients.
* Bridged clients pull histories from tox-homeserver, that is, all bridged clients hold a same local history. The *history* is the messages sent to tox-homeserver when bridged client is offline.
* tox-homeserver is still work in progress. Who interesting in can take a look at https://github.com/envsh/tox-homeserver.
* Since toxcore supports group for temporary for now, there is no (good) way to merge group messages while the group was recreated.

### Run a server

https://github.com/envsh/tox-homeserver/releases

Download the latest toxhs.tar.gz (for linux amd64)

    tar xvf toxhs.tar.gz   # file name is toxhs
    ./toxhs

tox-homeserver open 2 ports, `*:2080` (gRPC) and `*:8099` (Websocket + HTTP), for serve.

After `toxhs` is launched, you can see lines like below in terminal:

    2018-05-27 11:10:35.056859 I | [gofiat] xtox.go:68: ID: B5E7631D4D6C0EC0581B9DA34F431F0F2FAB8D115F9930D3DDA1333C8006A477A6440C224DD8

This is the server's (your's) toxid.

**Warning:** If you want to recompile tox-homeserver, you **have to** use the [envsh/go-toxcore-c](https://github.com/envsh/go-toxcore-c) instead of TokTok/go-toxcore-c. Because there were 2 important fixes haven't merged into [TokTok/go-toxcore-c](https://github.com/TokTok/go-toxcore-c). If you chhoose to use precompiled binary from [the release page][rel-page], you can ignore this warning safely.

### Run a native client

Goto [release page][rel-page] and download client binary suite for your environment:

| Platform         | Binary                          |
|:----------------:|:--------------------------------|
| Android ARM      | qofia-ffi-arm.apk               |
| Windowns x64     | qofia-ffi-amd64.tar.gz          |
| Windows x32      | qofia-ffi-i386.tar.gz           |
| Linux x64        | qofia-x86_64.AppImage           |
| Linux x64(tared) | qofia-linux-amd64-0.4.0.tar.bz2 |

To running on Windows or Linux(Except AppImage package), you should download suitable Qt runtime from https://github.com/qtchina/, put them into `bin` directory manually:

* https://github.com/qtchina/qtenv_win64
* https://github.com/qtchina/qtenv_win32
* https://github.com/qtchina/qt510_qt_inline

### Web client (demo)

Open browser and visit `http://your-ip-address:8099/webdui`.

### client snapshots

https://github.com/envsh/tox-homeserver/tree/master/snapshots

![Contact List](https://raw.githubusercontent.com/envsh/tox-homeserver/master/snapshots/contact_list.png)

![Message List](https://raw.githubusercontent.com/envsh/tox-homeserver/master/snapshots/message_list.png)

[rel-page]: https://github.com/envsh/tox-homeserver/releases/latest
