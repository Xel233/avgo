package avutil

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
	"image"
	"os"
	"image/jpeg"
	"image/png"
	"github.com/xel233/avgo"
)

func (this *AVFrame) cptr() *C.struct_AVFrame {
	return (*C.struct_AVFrame)(unsafe.Pointer(this))
}

func (this *AVFrame) getYCbCr420() (Y, Cb, Cr []byte) {
	f := this.cptr()
	w, h := int(f.linesize[0]), int(f.height)
	if d0 := unsafe.Pointer(f.data[0]); d0 != nil {
		Y = avgo.GoBytes(d0, w*h)
	}
	if d1 := unsafe.Pointer(f.data[1]); d1 != nil {
		Cb = avgo.GoBytes(d1, int(f.linesize[1])*h/2)
	}
	if d2 := unsafe.Pointer(f.data[2]); d2 != nil {
		Cr = avgo.GoBytes(d2, int(f.linesize[2])*h/2)
	}
	return
}

func (this *AVFrame) Free() {
	cptr := this.cptr()
	C.av_frame_free(&cptr)
}

func (this *AVFrame) IsKeyFrame() bool {
	return int(this.cptr().key_frame) > 0
}

func (this *AVFrame) GetImage() *ImageYCbCr {
	f := this.cptr()
	w, h := int(f.linesize[0]), int(f.height)
	r := image.Rectangle{image.Point{0, 0}, image.Point{w, h}}
	ycbcr := image.YCbCr{
		SubsampleRatio: image.YCbCrSubsampleRatio420,
		Rect: r,
		YStride: r.Dx(),
		CStride: (r.Max.X+1)/2 - r.Min.X/2,
	}
	ycbcr.Y, ycbcr.Cb, ycbcr.Cr = this.getYCbCr420()
	subR := image.Rectangle{
		image.Point{0, 0},
		image.Point{int(f.width),int(f.height)},
	}
	return (*ImageYCbCr)(ycbcr.SubImage(subR).(*image.YCbCr))
}

func (this *AVFrame) Data() [8]*C.uint8_t {
	return this.cptr().data
}

func (this *ImageYCbCr) SavePNG(pathToImage string) error {
	f, err := os.Create(pathToImage + ".png")
	if err == nil {
		defer f.Close()
		err = png.Encode(f, (*image.YCbCr)(this))
	}
	return err
}

func (this *ImageYCbCr) SaveJPEG(pathToImage string, quality int) error {
	f, err := os.Create(pathToImage + ".jpeg")
	if err == nil {
		defer f.Close()
		err = jpeg.Encode(f, (*image.YCbCr)(this), &jpeg.Options{quality})
	}
	return err
}
