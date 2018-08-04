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
import (
	"github.com/xel233/avgo/avutil"
	"unsafe"
)

func (this *AVCodecContext) cptr() *C.struct_AVCodecContext {
	return (*C.struct_AVCodecContext)(unsafe.Pointer(this))
}

func (this *AVCodecContext) Close() {
	C.avcodec_close(this.cptr())
}

func (this *AVCodecContext) CodecType() int32 {
	return this.cptr().codec_type
}

func (this *AVCodecContext) IsVideo() bool {
	return this.CodecType() == C.AVMEDIA_TYPE_VIDEO
}

func (this *AVCodecContext) CodecId() CodecId {
	return (CodecId)(this.cptr().codec_id)
}

func (this *AVCodecContext) AvcodecOpen2(codec *AVCodec, options **AVDictionary) int {
	return int(C.avcodec_open2(
		(*C.struct_AVCodecContext)(this),
		(*C.struct_AVCodec)(codec),
		(**C.struct_AVDictionary)(unsafe.Pointer(options)),
	))
}

func (this *AVCodecContext) Width() int {
	return int(this.cptr().width)
}

func (this *AVCodecContext) Height() int {
	return int(this.cptr().height)
}

func (this *AVCodecContext) AvcodecDecodeVideo2(picture *avutil.AVFrame, gotPicturePtr *int, avpkt *AVPacket) int {
	return int(C.avcodec_decode_video2(
		(*C.struct_AVCodecContext)(this),
		(*C.struct_AVFrame)(unsafe.Pointer(picture)),
		(*C.int)(unsafe.Pointer(gotPicturePtr)),
		(*C.struct_AVPacket)(avpkt),
	))
}
