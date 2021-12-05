package frameSetSaver

import (
	"bytes"
	"image/draw"
)

type SaverInterface interface {
	SetFramerate(framerate int) SaverInterface
	SetFrameSet(images *[]draw.Image) SaverInterface
	SaveInFile(path string) (output string, err error)
	SaveInBuffer() (output bytes.Buffer, err error)
}
