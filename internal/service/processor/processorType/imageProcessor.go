package processorType

import (
	"errors"
	"image"
	"image/draw"

	"github.com/disintegration/imaging"
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/mitchellh/mapstructure"
	"github.com/mo3golom/wonder-animator/internal/dto"
	"github.com/mo3golom/wonder-animator/internal/dto/processorOptions"
	"github.com/mo3golom/wonder-animator/pkg/loader"
)

type ImageProcessor struct {
	options *processorOptions.ImageOptions
	*ProcessorStruct
}

func (ip *ImageProcessor) Processing(_ *dto.Block, _ *dto.FrameData) (output *image.RGBA, err error) {
	if nil == ip.options {
		return nil, errors.New("предварительно необходимо преобразовать настройки")
	}

	img, err := loader.LoadImage(ip.options.Data)

	if nil != err {
		return nil, err
	}

	// Проверяем, нужен ли ресайз картинки
	originalSize := img.Bounds()
	width := originalSize.Dx()
	height := originalSize.Dy()
	isNeedResize := false

	// Проверяем нужно ли изменение ширины
	if 0 < ip.options.Width && width != ip.options.Width {
		width = ip.options.Width
		isNeedResize = true
	}

	// Проверяем нужно ли изменение высоты
	if 0 < ip.options.Height && height != ip.options.Height {
		height = ip.options.Height
		isNeedResize = true
	}

	// Делаем ресайз если нужно
	if isNeedResize {
		img = imaging.Resize(
			img,
			width,
			height,
			imaging.Lanczos,
		)
	}

	output = image.NewRGBA(img.Bounds())
	graphicContext := draw2dimg.NewGraphicContext(output)
	graphicContext.DrawImage(img)

	// Если в настройках есть макса, то рисуем маску
	if nil != ip.options.Mask {
		maskOptions := ip.options.Mask
		src, _ := loader.LoadImage(maskOptions.Src)
		mask, _ := loader.LoadImage(maskOptions.Mask)
		rect := image.Rect(
			maskOptions.X,
			maskOptions.Y,
			maskOptions.X+src.Bounds().Dx(),
			maskOptions.Y+src.Bounds().Dy(),
		)

		draw.DrawMask(output, rect, src, image.Pt(0, 0), mask, image.Pt(0, 0), draw.Over)
	}

	return output, nil
}

func (ip *ImageProcessor) TransformOptions(options *map[string]interface{}) ProcessorInterface {
	imageOptions := processorOptions.NewImageOptions()
	_ = mapstructure.Decode(*options, imageOptions)

	ip.options = imageOptions

	return ip
}
