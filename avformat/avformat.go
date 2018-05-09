package avformat

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
	"unsafe"
)

func AvRegisterAll() {
	C.av_register_all()
}

func AvformatOpenInput(ps **AVFormatContext, url string, fmt *AVInputFormat, options **AVDictionary) int {
	return int(C.avformat_open_input(
		(**C.struct_AVFormatContext)(unsafe.Pointer(ps)),
		C.CString(url),
		(*C.struct_AVInputFormat)(unsafe.Pointer(fmt)),
		(**C.struct_AVDictionary)(unsafe.Pointer(options)),
	))
}
