# avgo
**CGO binding for FFmpeg**

#### Installation

1. install ffmpeg
```bash
sudo apt-get -y --force-yes install autoconf automake build-essential libass-dev libfreetype6-dev libsdl1.2-dev libtheora-dev libtool libva-dev libvdpau-dev libvorbis-dev libxcb1-dev libxcb-shm0-dev libxcb-xfixes0-dev pkg-config texi2html zlib1g-dev
sudo apt-get install yasm
```
2. set env
```$xslt
export FFMPEG_ROOT=/usr/local/ffmpeg
export CGO_LDFLAGS="-L$FFMPEG_ROOT/lib/ -lavcodec -lavformat -lavutil -lswscale -lswresample -lavdevice -lavfilter"
export CGO_CFLAGS="-I$FFMPEG_ROOT/include"
export LD_LIBRARY_PATH=$FFMPEG_ROOT/lib
export PKG_CONFIG_PATH=$PKG_CONFIG_PATH:$FFMPEG_ROOT/lib/pkgconfig
```
3. go package
```bash
go get -u -v "github.com/xel233/avgo"
```

#### Usage

1. [example](https://github.com/Xel233/avgo/tree/master/example)

