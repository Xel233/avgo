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
	"github.com/xel233/avgo/avcodec"
	"unsafe"
)

func (this *AVFormatContext) cptr() *C.struct_AVFormatContext {
	return (*C.struct_AVFormatContext)(unsafe.Pointer(this))
}

func (this *AVFormatContext) Close() {
	C.avformat_close_input((**C.struct_AVFormatContext)(unsafe.Pointer(&this)))
}

func (this *AVFormatContext) Duration() int64 {
	return int64(this.cptr().duration)
}

func (this *AVFormatContext) AvformatFindStreamInfo(options **AVDictionary) int {
	return int(C.avformat_find_stream_info(
		this.cptr(),
		(**C.struct_AVDictionary)(unsafe.Pointer(&options)),
	))
}

func (this *AVFormatContext) NBStreams() int {
	return int(this.cptr().nb_streams)
}

func (this *AVFormatContext) Streams() []*AVStream {
	result := []*AVStream{}
	streams := (*AVStream)(unsafe.Pointer(this.cptr().streams))
	for i := 0; i < this.NBStreams(); i++ {
		if stream := streams.Select(i); stream != nil {
			result = append(result, stream)
		}
	}
	return result
}

func (this *AVFormatContext) AvReadFrame(pkt *avcodec.AVPacket) int {
	return int(C.av_read_frame((*C.struct_AVFormatContext)(unsafe.Pointer(this)), (*C.struct_AVPacket)(unsafe.Pointer(pkt))))
}
