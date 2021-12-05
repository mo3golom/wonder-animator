package processorType

import (
	"errors"
	"github.com/mo3golom/wonder-animator/internal/dto"
	"github.com/mo3golom/wonder-animator/internal/dto/processorOptions"
	"github.com/mo3golom/wonder-animator/internal/transformer"
	"github.com/mo3golom/wonder-animator/pkg/draw2dExtend"
	"image/draw"
	"strconv"
)

type TimerProcessor struct {
	options *processorOptions.TimerOptions
	*ProcessorStruct
}

func (tp *TimerProcessor) Processing(dest draw.Image, block *dto.Block, frameData *dto.FrameData) (err error) {
	if nil == tp.options {
		return errors.New("предварительно необходимо преобразовать настройки")
	}

	second := tp.calculateSecond(
		block.StartAt,
		block.Duration,
		frameData.Pos,
		frameData.FPS,
		tp.options.Mode(),
	)

	// Нужно ли включать 0 в таймер
	if !tp.options.IncludeZero() {
		second += 1
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
		DrawText(strconv.Itoa(second), tp.options.FontSize(), backgroundOptions)

	return nil
}

func (tp *TimerProcessor) TransformOptions(options *map[string]string) ProcessorInterface {
	tp.options = transformer.TransformTimerOptions(*options)

	return tp
}

func (tp *TimerProcessor) calculateSecond(startAt, duration float32, framePos int, framesPerSecond int, mode int) (result int) {
	switch mode {
	case processorOptions.TimerModeNormal:
		startFrame := transformer.SecondsToFrameCount(startAt, framesPerSecond)
		result = int(transformer.FrameCountToSeconds(framePos-startFrame, framesPerSecond))

	case processorOptions.TimerModeReverse:
		endFrame := transformer.SecondsToFrameCount(startAt+duration, framesPerSecond)
		result = int(transformer.FrameCountToSeconds(endFrame-framePos, framesPerSecond))
	}

	return result
}
