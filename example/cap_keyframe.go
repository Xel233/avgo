package main

import (
	"fmt"
	"github.com/xel233/avgo/avcodec"
	"github.com/xel233/avgo/avformat"
	"github.com/xel233/avgo/avutil"
	"time"
)

func main() {
	start := time.Now().Unix()
	avformat.AvRegisterAll()
	fi := "sample.mp4"
	pFormatCtx, ret := avformat.AvformatOpenInput2(fi, nil, nil)
	if ret < 0 {
		panic("can not open file")
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
	i := 1
	packet := avcodec.AvInitPakcet()
	var frameFinished int
	for int(pFormatCtx.AvReadFrame(packet)) >= 0 {
		if packet.StreamIndex() == videoStream {
			pCodecCtx.AvcodecDecodeVideo2(pFrame, &frameFinished, packet)
			if frameFinished > 0 {
				if pFrame.IsKeyFrame() {
					filePath := fmt.Sprintf("keyframe_%d", packet.PTS())
					fmt.Printf("[%v]Captured %s.jpeg\n", i, filePath)
					go pFrame.GetImage().SaveJPEG(filePath, 50)
					i++
				}
			}
		}
		packet.Free()
	}
	fmt.Printf("Timecosts：%ds\n", time.Now().Unix()-start)
}
