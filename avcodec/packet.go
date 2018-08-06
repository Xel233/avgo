package avcodec

//#cgo pkg-config: libavformat libavcodec libavutil libavdevice libavfilter libswresample libswscale
//#include <stdio.h>
//#include <stdlib.h>
//#include <inttypes.h>
//#include <stdint.h>
//#include <string.h>
//#include <libavformat/avformat.h>
//#include <libavcodec/avcodec.h>
//#include <libavutil/avutil.h>
//#include <libavutil/opt.h>
//#include <libavdevice/avdevice.h>
import "C"

func (this *AVPacket) cptr() *C.struct_AVPacket {
	return (*C.struct_AVPacket)(this)
}

func (this *AVPacket) Free() {
	C.av_free_packet(this.cptr())
}

func (this *AVPacket) StreamIndex() int {
	return int(this.cptr().stream_index)
}

func (this *AVPacket) PTS() int64 {
	return int64(this.cptr().pts)
}

func (this *AVPacket) DTS() int64 {
	return int64(this.cptr().dts)
}
