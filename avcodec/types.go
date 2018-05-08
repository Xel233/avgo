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

type (
	AVCodecContext C.struct_AVCodecContext
	AVCodec        C.struct_AVCodec
	AVDictionary   C.struct_AVDictionary
	AVPacket       C.struct_AVPacket
	CodecId        C.enum_AVCodecID
)
