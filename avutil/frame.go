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
)

func (this *AVFrame) cptr() *C.struct_AVFrame {
	return (*C.struct_AVFrame)(unsafe.Pointer(this))
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
	img := image.NewYCbCr(r, image.YCbCrSubsampleRatio420)
	img.Y = C.GoBytes(unsafe.Pointer(f.data[0]), C.int(w*h))
	wCb := int(f.linesize[1])
	if unsafe.Pointer(f.data[1]) != nil {
		img.Cb = C.GoBytes(unsafe.Pointer(f.data[1]), C.int(wCb*h/2))
	}
	wCr := int(f.linesize[2])
	if unsafe.Pointer(f.data[2]) != nil {
		img.Cr = C.GoBytes(unsafe.Pointer(f.data[2]), C.int(wCr*h/2))
	}
	return (*ImageYCbCr)(img)
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