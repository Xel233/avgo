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
	"github.com/xel233/avgo/avcodec"
)

func (this *AVStream) cptr() **C.struct_AVStream {
	return (**C.struct_AVStream)(unsafe.Pointer(this))
}

func (this *AVStream) Select(idx int) *AVStream {
	offset1 := (unsafe.Sizeof(unsafe.Pointer(*this.cptr())) * uintptr(idx))
	stream := *(**C.struct_AVStream)(unsafe.Pointer(uintptr(unsafe.Pointer(this.cptr())) + offset1))
	return (*AVStream)(stream)
}

func (this *AVStream) Codec() *avcodec.AVCodecContext {
	c := (*C.struct_AVStream)(unsafe.Pointer(this)).codec
	return (*avcodec.AVCodecContext)(unsafe.Pointer(c))
}