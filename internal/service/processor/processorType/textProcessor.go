package processorType

import (
	"errors"
	"github.com/mo3golom/wonder-animator/internal/dto"
	"github.com/mo3golom/wonder-animator/internal/dto/processorOptions"
	"github.com/mo3golom/wonder-animator/internal/transformer"
	"github.com/mo3golom/wonder-animator/pkg/draw2dExtend"
	"image/draw"
)

type TextProcessor struct {
	options *processorOptions.TextOptions
	*ProcessorStruct
}

func (tp *TextProcessor) Processing(dest draw.Image, block *dto.Block, frameData *dto.FrameData) (err error) {
	if nil == tp.options {
		return errors.New("предварительно необходимо преобразовать настройки")
	}

	if "" == tp.options.Text() {
		return errors.New("для генерации текста нужен текст")
	}

	graphicContext := draw2dExtend.NewGraphicContext(dest)
	effectValues := tp.applyEffects(block, frameData)

	backgroundOptions := draw2dExtend.NewBackgroundOptions().
		SetFillColor(tp.options.BackgroundColor()).
		SetPadding(tp.options.Padding()).
		SetRadius(tp.options.Radius())

	tp.drawBuilder.
		SetGraphicContext(graphicContext).
		SetX(effectValues.X()).
		SetY(effectValues.Y()).
		SetRotate(effectValues.Rotate()).
		SetOpacity(effectValues.Opacity()).
		SetScale(effectValues.Scale()).
		SetRotatePointType(effectValues.RotatePoint).
		SetFillColor(tp.options.TextColor()).
		DrawText(tp.options.Text(), tp.options.FontSize(), backgroundOptions)

	return nil
}

func (tp *TextProcessor) TransformOptions(options *map[string]string) ProcessorInterface {
	tp.options = transformer.TransformTextOptions(*options)

	return tp
}
