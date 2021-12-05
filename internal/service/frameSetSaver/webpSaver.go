package frameSetSaver

import (
	"bytes"
	"github.com/mo3golom/wonder-animator/internal/transformer"
	"github.com/sizeofint/webpanimation"
	"image/draw"
	"os"
)

type WebpSaver struct {
	images    *[]draw.Image
	frameRate int
}

func NewWebpSaver() *WebpSaver {
	return &WebpSaver{}
}

func (g *WebpSaver) SetFramerate(framerate int) SaverInterface {
	g.frameRate = framerate

	return g
}

func (g *WebpSaver) SetFrameSet(images *[]draw.Image) SaverInterface {
	g.images = images

	return g
}

func (g *WebpSaver) SaveInFile(path string) (output string, err error) {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0600)

	if err != nil {
		return
	}

	defer f.Close()

	buffer, err := g.SaveInBuffer()

	if err != nil {
		return
	}

	_, err = f.Write(buffer.Bytes())

	if err != nil {
		return
	}

	return path, nil
}

func (g *WebpSaver) SaveInBuffer() (output bytes.Buffer, err error) {
	images := *g.images
	rectangle := images[0].Bounds()

	webpAnimation := webpanimation.NewWebpAnimation(rectangle.Dx(), rectangle.Dy(), 0)
	defer webpAnimation.ReleaseMemory()

	// Настройки взяты из примеров библиотеки
	webpConfig := webpanimation.NewWebpConfig()
	webpConfig.SetLossless(1) // Устанавливаем качество

	// Высчитываем сколько должен длиться один кадр
	// Сначала находим продолжительность видео в миллисекундах
	duration := transformer.FrameCountToSeconds(len(images), g.frameRate) * 1000
	// Затем делим продолжтиельность на число кадров
	// В данном случае мы получим продолжительность гифки максимально близкую к требуемой (но не факт что равной)
	frameDisplayTime := int(duration) / len(images)
	timeline := 0

	for _, img := range images {
		_ = webpAnimation.AddFrame(img, timeline, webpConfig)
		timeline += frameDisplayTime
	}

	err = webpAnimation.Encode(&output)

	return output, err
}
