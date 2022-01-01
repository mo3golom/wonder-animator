package processorType

import (
	"errors"
	"image"
	"math"
	"time"

	"github.com/mitchellh/mapstructure"
	"github.com/mo3golom/wonder-animator/internal/dto"
	"github.com/mo3golom/wonder-animator/internal/dto/processorOptions"
	"github.com/mo3golom/wonder-animator/internal/transformer"
)

type TimerProcessor struct {
	options *processorOptions.TimerOptions
	*TextProcessor
}

func (tp *TimerProcessor) Processing(block *dto.Block, frameData *dto.FrameData) (output *image.RGBA, err error) {
	if nil == tp.options {
		return nil, errors.New("предварительно необходимо преобразовать настройки")
	}

	second := 0.0

	switch tp.options.Mode {
	case processorOptions.TimerModeNormal:
		startFrame := transformer.SecondsToFrameCount(float64(block.StartAt), frameData.FPS)
		second = transformer.FrameCountToSeconds(frameData.Pos-startFrame, frameData.FPS)

	case processorOptions.TimerModeReverse:
		endFrame := transformer.SecondsToFrameCount(float64(block.StartAt+block.Duration), frameData.FPS)
		second = transformer.FrameCountToSeconds(endFrame-frameData.Pos, frameData.FPS)
	}

	// Нужно ли включать 0 в таймер
	if !tp.options.IncludeZero {
		second += 1
	}

	// Преобразуем вычисленную секунду в тип Time
	// и выводим время в форматированном формате
	// Для написания правильной строки форматирования см. https://yourbasic.org/golang/format-parse-string-time-date-example/
	sec, dec := math.Modf(second)
	tp.TextProcessor.options.Text = time.
		Unix(int64(sec), int64(dec*(1e9))).
		Format(tp.options.Format)

	return tp.TextProcessor.Processing(block, frameData)
}

func (tp *TimerProcessor) TransformOptions(options *map[string]interface{}) ProcessorInterface {
	timerOptions := processorOptions.NewTimerOptions()
	_ = mapstructure.Decode(*options, timerOptions)

	tp.options = timerOptions

	// Преобразуем настройки для процессора текста
	tp.TextProcessor.TransformOptions(options)

	return tp
}
