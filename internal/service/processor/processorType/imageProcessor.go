package processorType

import (
	"errors"
	"github.com/disintegration/imaging"
	"github.com/mo3golom/wonder-animator/internal/dto"
	"github.com/mo3golom/wonder-animator/internal/dto/processorOptions"
	"github.com/mo3golom/wonder-animator/internal/transformer"
	"github.com/mo3golom/wonder-animator/pkg/draw2dExtend"
	"github.com/mo3golom/wonder-animator/pkg/imagingExtend"
	"github.com/mo3golom/wonder-animator/pkg/loader"
	"image/draw"
	"math"
)

type ImageProcessor struct {
	options *processorOptions.ImageOptions
	*ProcessorStruct
}

func (ip *ImageProcessor) Processing(dest draw.Image, block *dto.Block, frameData *dto.FrameData) (err error) {
	if nil == ip.options {
		return errors.New("предварительно необходимо преобразовать настройки")
	}

	img, err := loader.LoadImage(ip.options.Data())

	if nil != err {
		return err
	}

	effectValues := ip.applyEffects(block, frameData)

	// Проверяем, нужен ли ресайз картинки
	originalSize := img.Bounds()
	width := originalSize.Dx()
	height := originalSize.Dy()
	isNeedResize := false

	// Проверяем нужно ли изменение ширины
	if 0 < ip.options.Width() && width != ip.options.Width() {
		width = ip.options.Width()
		isNeedResize = true
	}

	// Проверяем нужно ли изменение высоты
	if 0 < ip.options.Height() && height != ip.options.Height() {
		height = ip.options.Height()
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

	// Устанавливаем прозрачность
	img = imagingExtend.Opacity(img, float64(effectValues.Opacity()))

	// Делаем поворот
	rotatePoint := draw2dExtend.GetRotatePointByType(
		effectValues.RotatePoint,
		effectValues.X(),
		effectValues.Y(),
		float64(width),
		float64(height),
	)

	img = imagingExtend.RotateAround(
		img,
		effectValues.Rotate(),
		math.Abs(effectValues.X()-rotatePoint.X),
		math.Abs(effectValues.Y()-rotatePoint.Y),
	)

	graphicContext := draw2dExtend.NewGraphicContext(dest)
	graphicContext.Scale(effectValues.Scale(), rotatePoint)
	graphicContext.Translate(effectValues.X(), effectValues.Y())
	graphicContext.DrawImage(img)

	return nil
}

func (ip *ImageProcessor) TransformOptions(options *map[string]string) ProcessorInterface {
	ip.options = transformer.TransformImageOptions(*options)

	return ip
}
