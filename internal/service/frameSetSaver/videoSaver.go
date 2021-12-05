package frameSetSaver

import (
	"bytes"
	"github.com/mo3golom/wonder-animator/pkg/ffmpeg"
	"image/draw"
)

type VideoSaver struct {
	ffmpeg *ffmpeg.FFMpeg
	images *[]draw.Image
}

func NewVideoSaverStrategy(ffmpeg *ffmpeg.FFMpeg) *VideoSaver {
	return &VideoSaver{ffmpeg: ffmpeg}
}

func (v *VideoSaver) SetFramerate(framerate int) SaverInterface {
	v.ffmpeg.SetFramerate(framerate)

	return v
}

func (v *VideoSaver) SetFrameSet(images *[]draw.Image) SaverInterface {
	v.images = images

	return v
}

func (v *VideoSaver) SaveInFile(path string) (output string, err error) {
	err = v.ffmpeg.SetOutput(path).CreateVideoFromImages(*v.images)

	return path, err
}

func (v *VideoSaver) SaveInBuffer() (output bytes.Buffer, err error) {
	panic("video save dont support save in buffer")
}
