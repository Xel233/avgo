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
import "unsafe"

func AvcodecFindDecoder(id CodecId) *AVCodec {
	dec := C.avcodec_find_decoder((C.enum_AVCodecID)(id))
	return (*AVCodec)(unsafe.Pointer(dec))
}

func AvInitPakcet() *AVPacket {
	avpkt := &C.struct_AVPacket{}
	C.av_init_packet(avpkt)
	return (*AVPacket)(unsafe.Pointer(avpkt))
}
