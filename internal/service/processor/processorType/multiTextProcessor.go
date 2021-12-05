package processorType

import (
	"errors"
	"github.com/mo3golom/wonder-animator/internal/dto"
	"github.com/mo3golom/wonder-animator/internal/dto/processorOptions"
	"github.com/mo3golom/wonder-animator/internal/transformer"
	"github.com/mo3golom/wonder-animator/pkg/draw2dExtend"
	"github.com/mo3golom/wonder-animator/pkg/strings"
	"image/draw"
)

type MultiTextProcessor struct {
	options *processorOptions.MultiTextOptions
	*ProcessorStruct
}

func (mtp *MultiTextProcessor) Processing(dest draw.Image, block *dto.Block, frameData *dto.FrameData) (err error) {
	var wordsMaxAvailableAppend float64 = 0
	var wordsAppend float64 = 0
	var totalMaxLineHeight float64 = 0

	if nil == mtp.options {
		return errors.New("предварительно необходимо преобразовать настройки")
	}

	if "" == mtp.options.Text() {
		return errors.New("для генерации текста нужен текст")
	}

	// Разделяем текст на слова
	lines := strings.WordWrap(mtp.options.Text(), mtp.options.LineWidth())

	// Вычисляем максимальный кадр, до которого эффект вывода еще будет применяться
	// После максимального кадра будут выводиться все слова
	frameMax := frameData.Max - (frameData.FPS * 2)

	if 0 > frameMax {
		frameMax = frameData.FPS
	}

	totalWords := 0

	for _, words := range lines {
		totalWords += len(words)
	}

	// Применяем режим вывода
	switch mtp.options.Mode() {
	case processorOptions.MultiTextModeAll:
		wordsMaxAvailableAppend = float64(totalWords)

	case processorOptions.MultiTextModeAfter:
		wordsMaxAvailableAppend = (float64(totalWords) * float64(frameData.Pos)) / float64(frameMax)
	}

	wordStep := frameMax / totalWords
	backgroundOptions := draw2dExtend.NewBackgroundOptions().
		SetFillColor(mtp.options.BackgroundColor()).
		SetPadding(mtp.options.Padding()).
		SetRadius(mtp.options.Radius())

	for _, words := range lines {
		var totalLineWidth float64 = 0
		var maxLineHeight float64 = 0

		for _, word := range words {
			corrector := int(wordsAppend) * wordStep
			progress := float32(frameData.Pos-corrector) / float32(frameMax-corrector)

			if 1 < progress {
				progress = 1
			}

			effectValues := mtp.applyEffectsWidthAdditionalData(block, progress, totalLineWidth, totalMaxLineHeight)

			width, height := mtp.drawBuilder.
				SetGraphicContext(draw2dExtend.NewGraphicContext(dest)).
				SetX(effectValues.X()).
				SetY(effectValues.Y()).
				SetRotate(effectValues.Rotate()).
				SetOpacity(effectValues.Opacity()).
				SetScale(effectValues.Scale()).
				SetRotatePointType(effectValues.RotatePoint).
				SetFillColor(mtp.options.TextColor()).
				DrawText(word, mtp.options.FontSize(), backgroundOptions)

			// Добавляем длину слова к длине строки
			totalLineWidth += width + mtp.options.Margin()

			// Находим максимальную высоту строки
			if height > maxLineHeight {
				maxLineHeight = height + mtp.options.Margin()
			}

			wordsAppend++

			// Если добавили все слова по количеству, то завершаем
			if wordsAppend >= wordsMaxAvailableAppend {
				return nil
			}
		}

		// Добавляем максимальную высоту строки к высоте всего текста
		totalMaxLineHeight += maxLineHeight
	}

	return nil
}

func (mtp *MultiTextProcessor) TransformOptions(options *map[string]string) ProcessorInterface {
	mtp.options = transformer.TransformMultiTextOptions(*options)

	return mtp
}
