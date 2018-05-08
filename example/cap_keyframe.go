package main

import (
	"fmt"
	"time"
	"github.com/xel233/avgo/avformat"
	"github.com/xel233/avgo/avcodec"
	"github.com/xel233/avgo/avutil"
)

func main() {
	start := time.Now().Unix()
	avformat.AvRegisterAll()
	fi := "sample.mp4"
	pFormatCtx, err := avformat.AvformatOpenInput(fi, nil, nil)
	if err != nil {
		panic(err.Error())
	}
	defer pFormatCtx.Close()
	if pFormatCtx.AvformatFindStreamInfo(nil) < 0 {
		panic("can not find stream info")
	}
	var pCodecCtx *avcodec.AVCodecContext
	videoStream := -1
	for idx, stream := range pFormatCtx.Streams() {
		pCodecCtx = stream.Codec()
		if pCodecCtx.IsVideo() {
			videoStream = idx
			break
		}
	}
	if videoStream == -1 {
		panic("can not find video stream")
	}
	defer pCodecCtx.Close()
	pCodec := avcodec.AvcodecFindDecoder(pCodecCtx.CodecId())
	if pCodec == nil {
		panic("can not find codec")
	}
	if pCodecCtx.AvcodecOpen2(pCodec, nil) < 0 {
		panic("can not open codec")
	}
	pFrame := avutil.AvFrameAlloc()
	defer pFrame.Free()
	i := 0
	var packet avcodec.AVPacket
	var frameFinished int
	for int(pFormatCtx.AvReadFrame(&packet)) >= 0 {
		if packet.StreamIndex() == videoStream {
			pCodecCtx.AvcodecDecodeVideo2(pFrame, &frameFinished, &packet)
			if frameFinished > 0 {
				if pFrame.IsKeyFrame() {
					filePath := fmt.Sprintf("keyframe_%d", i)
					fmt.Printf("[%v]Captured %s\n", packet.PTS(), filePath)
					go pFrame.GetImage().SaveJPEG(filePath, 50)
					i++
				}
			}
		}
		packet.Free()
	}
	/*
	C.av_register_all()
	var pFormatCtx *C.struct_AVFormatContext
	fi := "sample.mp4"
	if int(C.avformat_open_input(&pFormatCtx, C.CString(fi), nil, nil)) < 0 {
		panic("can not find stream info")
	}
	if int(C.avformat_find_stream_info(pFormatCtx, nil)) < 0 {
		panic("can not find video stream")
	}
	//C.av_dump_format(pFormatCtx, 0, C.CString(""), 0)
	var pCodecCtx *C.struct_AVCodecContext
	var streams *C.struct_AVStream
	videoStream := -1
	for i := 0; i < int(pFormatCtx.nb_streams); i++ {
		offset := (unsafe.Sizeof(unsafe.Pointer(*pFormatCtx.streams)) * uintptr(i))
		streams = *(**C.struct_AVStream)(unsafe.Pointer(uintptr(unsafe.Pointer(pFormatCtx.streams)) + offset))
		pCodecCtx = (*C.struct_AVCodecContext)(streams.codec)
		if pCodecCtx.codec_type == C.AVMEDIA_TYPE_VIDEO {
			videoStream = i
			break
		}
	}
	if videoStream == -1 {
		panic("can not find video stream")
	}
	var pCodec *C.struct_AVCodec = C.avcodec_find_decoder(pCodecCtx.codec_id)
	if pCodec == nil {
		panic("can not find codec")
	}
	if int(C.avcodec_open2(pCodecCtx, pCodec, nil)) < 0 {
		panic("can not open codec")
	}
	pFrame := C.av_frame_alloc()
	i := 0
	var packet C.struct_AVPacket
	for int(C.av_read_frame(pFormatCtx, &packet)) >= 0 {
		if int(packet.stream_index) == videoStream {
			var frameFinished C.int
			C.avcodec_decode_video2(pCodecCtx, pFrame, &frameFinished, &packet)
			if frameFinished > 0 {
				if pFrame.key_frame == 1 {
					img := getImage(pFrame)
					go saveFrame(img, 60, i)
					i++
				}
			}
		}
		C.av_free_packet(&packet)
	}
	C.av_frame_free(&pFrame)
	C.avcodec_close(pCodecCtx)
	C.avformat_close_input(&pFormatCtx)*/
	fmt.Printf("Timecosts：%ds\n", time.Now().Unix()-start)
}
