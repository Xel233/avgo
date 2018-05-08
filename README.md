# avgo

#### 项目介绍
**CGO实现的ffmpeg封装**

#### 安装教程

1. 安装ffmpeg
```bash
sudo apt-get -y --force-yes install autoconf automake build-essential libass-dev libfreetype6-dev libsdl1.2-dev libtheora-dev libtool libva-dev libvdpau-dev libvorbis-dev libxcb1-dev libxcb-shm0-dev libxcb-xfixes0-dev pkg-config texi2html zlib1g-dev
sudo apt-get install yasm
```
2. 设置环境变量
```$xslt
export FFMPEG_ROOT=/usr/local/ffmpeg
export CGO_LDFLAGS="-L$FFMPEG_ROOT/lib/ -lavcodec -lavformat -lavutil -lswscale -lswresample -lavdevice -lavfilter"
export CGO_CFLAGS="-I$FFMPEG_ROOT/include"
export LD_LIBRARY_PATH=$FFMPEG_ROOT/ffmpeg/lib
export PKG_CONFIG_PATH=$PKG_CONFIG_PATH:$FFMPEG_ROOT/lib/pkgconfig
```
3. 安装package
```bash
go get -u -v "github.com/xel233/avgo"
```

#### 使用说明

1. [example](https://github.com/Xel233/avgo/tree/master/example)

#### 参与贡献

1. Fork 本项目
2. 新建 Feat_xxx 分支
3. 提交代码
4. 新建 Pull Request
