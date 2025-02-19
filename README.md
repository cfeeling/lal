
# LAL

[![Release](https://img.shields.io/github/tag/q191201771/lal.svg?label=release)](https://github.com/cfeeling/lal/releases)
[![TravisCI](https://www.travis-ci.org/q191201771/lal.svg?branch=master)](https://www.travis-ci.org/q191201771/lal)
[![goreportcard](https://goreportcard.com/badge/github.com/cfeeling/lal)](https://goreportcard.com/report/github.com/cfeeling/lal)
![wechat](https://img.shields.io/:微信-q191201771-blue.svg)
![qqgroup](https://img.shields.io/:QQ群-1090510973-blue.svg)

[中文文档](https://pengrl.com/lal/#/)

LAL is an audio/video live streaming broadcast server written in Go. It's sort of like `nginx-rtmp-module`, but easier to use and with more features, e.g RTMP, RTSP(RTP/RTCP), HLS, HTTP[S]/WebSocket[s]-FLV/TS, H264/H265/AAC, relay, cluster, record, HTTP API/Notify, GOP cache.

And [more than a server, act as package and client](https://github.com/cfeeling/lal#more-than-a-server-act-as-package-and-client)

## Install

There are 2 ways of installing lal.

### Prebuilt binaries

Prebuilt binaries for Linux, macOS(Darwin), Windows are available in the [lal github releases page](https://github.com/cfeeling/lal/releases). Naturally, using [the latest release binary](https://github.com/cfeeling/lal/releases/latest) is the recommended way. The naming format is `lal_<version>_<platform>.zip`, e.g. `lal_v0.20.0_linux.zip`

LAL could also be built from the source wherever the Go compiler toolchain can run, e.g. for other architectures including arm32 and mipsle which have been tested by the community.

### Building from source

First, make sure that Go version >= 1.13

For Linux/macOS user:

```shell
$git clone https://github.com/cfeeling/lal.git
$cd lal
$make build
```

Then all binaries go into the `./bin/` directory. That's it.

For an experienced gopher(and Windows user), the only thing you should be concern is that `the main function` is under the `./app/lalserver` directory. So you can also:

```shell
$git clone https://github.com/cfeeling/lal.git
$cd lal/app/lalserver
$go build
```

Or using whatever IDEs you'd like.

So far, the only direct and indirect **dependency** of lal is [naza(A basic Go utility library)](https://github.com/cfeeling/lal.git) which is also written by myself. This leads to less dependency or version manager issues.

## Using

Running lalserver:

```
$./bin/lalserver -c ./conf/lalserver.conf.json
```

Using whatever clients you are familiar with to interact with lalserver.

For instance, publish rtmp stream to lalserver via ffmpeg:

```shell
$ffmpeg -re -i demo.flv -c:a copy -c:v copy -f flv rtmp://127.0.0.1:1935/live/test110
```

Play multi protocol stream from lalserver via ffplay:

```shell
$ffplay rtmp://127.0.0.1/live/test110
$ffplay http://127.0.0.1:8080/live/test110.flv
$ffplay http://127.0.0.1:8080/hls/test110/playlist.m3u8
$ffplay http://127.0.0.1:8080/hls/test110/record.m3u8
$ffplay http://127.0.0.1:8080/live/test110.ts
```

## More than a server, act as package and client

Besides a live stream broadcast server which named `lalserver` precisely, `project lal` even provides many other applications, e.g. push/pull/remux stream clients, bench tools, examples. Each subdirectory under the `./app/demo` directory represents a tiny demo.

Our goals are not only a production server but also a simple package with a well-defined, user-facing API, so that users can build their own applications on it.

`LAL` stands for `Live And Live` if you may wonder.


## Contact

Bugs, questions, suggestions, anything related or not, feel free to contact me with [lal github issues](https://github.com/cfeeling/lal/issues).

## License

MIT, see [License](https://github.com/cfeeling/lal/blob/master/LICENSE).
