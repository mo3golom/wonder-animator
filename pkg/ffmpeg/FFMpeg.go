package ffmpeg

import (
	"bytes"
	"image/draw"
	"image/png"
	"os/exec"
)

const (
	commandName = "ffmpeg"
	stdIn       = "pipe:0"
)

type FFMpeg struct {
	input, framerate, vcodec, output string
}

func (F *FFMpeg) SetInput(input string) *FFMpeg {
	F.input = input

	return F
}

func (F *FFMpeg) SetFramerate(framerate int) *FFMpeg {
	F.framerate = string(rune(framerate))

	return F
}

func (F *FFMpeg) SetVcodec(vcodec string) *FFMpeg {
	F.vcodec = vcodec

	return F
}

func (F *FFMpeg) SetOutput(output string) *FFMpeg {
	F.output = output

	return F
}

func NewFFMpeg() *FFMpeg {
	return &FFMpeg{vcodec: "png", framerate: "24"}
}

func (F *FFMpeg) CreateVideoFromImages(images []draw.Image) (err error) {
	cmd := F.SetInput(stdIn).createCommand()

	stdin, err := cmd.StdinPipe() // Open stdin pipe

	if nil != err {
		return
	}

	err = cmd.Start() // Start a process on another goroutine

	if nil != err {
		return
	}

	for _, image := range images {
		buf := new(bytes.Buffer)
		err = png.Encode(buf, image)

		if nil != err {
			continue
		}

		_, _ = stdin.Write(buf.Bytes())
	}

	err = stdin.Close()

	if nil != err {
		return
	}

	err = cmd.Wait()

	if nil != err {
		return
	}

	return nil
}

func (F *FFMpeg) createCommand() *exec.Cmd {
	return exec.Command(commandName, "-y", // Yes to all
		"-i", F.input, // take stdin as input
		"-framerate", F.framerate, // suppress "Frame rate very high for a muxer not efficiently supporting it"
		"-vcodec", F.vcodec, // Down sample audio birate to 128k
		F.output, // output to stdout
	)
}
