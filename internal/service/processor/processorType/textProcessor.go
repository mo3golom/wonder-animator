package processorType

import (
	"errors"
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/mitchellh/mapstructure"
	"github.com/mo3golom/wonder-animator/internal/dto"
	"github.com/mo3golom/wonder-animator/internal/dto/processorOptions"
	"github.com/mo3golom/wonder-animator/pkg/draw2dExtend"
	"image"
)

type TextProcessor struct {
	options *processorOptions.TextOptions
	*ProcessorStruct
}

func (tp *TextProcessor) Processing(_ *dto.Block, _ *dto.FrameData) (output *image.RGBA, err error) {
	if nil == tp.options {
		return nil, errors.New("предварительно необходимо преобразовать настройки")
	}

	if "" == tp.options.Text {
		return nil, errors.New("для генерации текста нужен текст")
	}

	output = image.NewRGBA(image.Rect(0, 0, tp.options.Width, tp.options.Height))
	graphicContext := draw2dimg.NewGraphicContext(output)

	backgroundOptions := draw2dExtend.NewBackgroundOptions()
	backgroundOptions.FillColor = draw2dExtend.ParseHexColor(tp.options.BackgroundColor)
	backgroundOptions.Padding = tp.options.Padding
	backgroundOptions.Radius = tp.options.Radius

	tp.DrawBuilder.
		SetGraphicContext(graphicContext).
		SetFillColor(draw2dExtend.ParseHexColor(tp.options.TextColor)).
		DrawText(tp.options.Text, tp.options.FontSize, backgroundOptions)

	return output, nil
}

func (tp *TextProcessor) TransformOptions(options *map[string]interface{}) ProcessorInterface {
	textOptions := processorOptions.NewTextOptions()
	_ = mapstructure.Decode(*options, textOptions)

	tp.options = textOptions

	return tp
}
