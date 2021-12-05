package transformer

import (
	"github.com/mo3golom/wonder-animator/internal/dto/processorOptions"
	"strconv"
)

func TransformTextOptions(options map[string]string) *processorOptions.TextOptions {
	textOptions := processorOptions.NewTextOptions()

	// Плучаем текст из настроек
	text, keyExists := options["text"]

	if keyExists {
		textOptions.SetText(text)
	}

	// Получаем радиус скругления из настроек
	fontSize, err := strconv.ParseFloat(options["fontSize"], 64)

	if nil == err {
		textOptions.SetFontSize(fontSize)
	}

	// Получаем отступ из настроек
	padding, err := strconv.ParseFloat(options["padding"], 64)

	if nil == err {
		textOptions.SetPadding(padding)
	}

	// Получаем радиус скругления из настроек
	radius, err := strconv.ParseFloat(options["radius"], 64)

	if nil == err {
		textOptions.SetRadius(radius)
	}

	// Получаем цвет текста
	textColor, keyExists := options["textColor"]

	if keyExists {
		textOptions.SetTextColor(textColor)
	}

	// Получаем цвет текста
	backgroundColor, keyExists := options["backgroundColor"]

	if keyExists {
		textOptions.SetBackgroundColor(backgroundColor)
	}

	return textOptions
}

func TransformTimerOptions(options map[string]string) *processorOptions.TimerOptions {
	timerOptions := processorOptions.NewTimerOptions(
		TransformTextOptions(options),
	)

	// Получаем режим таймера (нормальный или обратный)
	mode, err := strconv.Atoi(options["mode"])

	if nil == err {
		timerOptions.SetMode(mode)
	}

	// Нужно ли включать 0 в таймер
	includeZero, err := strconv.ParseBool(options["includeZero"])

	if nil == err {
		timerOptions.SetIncludeZero(includeZero)
	}

	return timerOptions
}

func TransformRectangleOptions(options map[string]string) *processorOptions.RectangleOptions {
	rectangleOptions := processorOptions.NewRectangleOptions()

	// Получаем ширину
	width, err := strconv.ParseFloat(options["width"], 64)

	if nil == err {
		rectangleOptions.SetWidth(width)
	}

	// Получаем высоту
	height, err := strconv.ParseFloat(options["height"], 64)

	if nil == err {
		rectangleOptions.SetHeight(height)
	}

	// Получаем закругление краев
	roundedRadius, err := strconv.ParseFloat(options["roundedRadius"], 64)

	if nil == err {
		rectangleOptions.SetRoundedRadius(roundedRadius)
	}

	// Получаем ширину обводки
	strokeWidth, err := strconv.ParseFloat(options["strokeWidth"], 64)

	if nil == err {
		rectangleOptions.SetStrokeWidth(strokeWidth)
	}

	// Получаем цвет заполнения
	fillColor, keyExists := options["fillColor"]

	if keyExists {
		rectangleOptions.SetFillColor(fillColor)
	}

	// Получаем цвет обводки
	strokeColor, keyExists := options["strokeColor"]

	if keyExists {
		rectangleOptions.SetStrokeColor(strokeColor)
	}

	return rectangleOptions
}

func TransformCircleOptions(options map[string]string) *processorOptions.CircleOptions {
	circleOptions := processorOptions.NewCircleOptions()

	// Получаем радиус
	radius, err := strconv.ParseFloat(options["radius"], 64)

	if nil == err {
		circleOptions.SetRadius(radius)
	}

	// Получаем ширину обводки
	strokeWidth, err := strconv.ParseFloat(options["strokeWidth"], 64)

	if nil == err {
		circleOptions.SetStrokeWidth(strokeWidth)
	}

	// Получаем цвет заполнения
	fillColor, keyExists := options["fillColor"]

	if keyExists {
		circleOptions.SetFillColor(fillColor)
	}

	// Получаем цвет обводки
	strokeColor, keyExists := options["strokeColor"]

	if keyExists {
		circleOptions.SetStrokeColor(strokeColor)
	}

	return circleOptions
}

func TransformImageOptions(options map[string]string) *processorOptions.ImageOptions {
	imageOptions := processorOptions.NewImageOptions()

	// Получаем данные картинки (base64, url, относительный путь)
	data, keyExists := options["data"]

	if keyExists {
		imageOptions.SetData(data)
	}

	// Получаем ширину
	width, err := strconv.Atoi(options["width"])

	if nil == err {
		imageOptions.SetWidth(width)
	}

	// Получаем высоту
	height, err := strconv.Atoi(options["height"])

	if nil == err {
		imageOptions.SetHeight(height)
	}

	return imageOptions
}

func TransformMultiTextOptions(options map[string]string) *processorOptions.MultiTextOptions {
	multiTextOptions := processorOptions.NewMultiTextOptions(
		TransformTextOptions(options),
	)

	// Получаем режим таймера (вывести все сразу или последовательно)
	mode, ok := options["mode"]

	if ok {
		multiTextOptions.SetMode(mode)
	}

	// Получаем отступ между словами
	margin, err := strconv.ParseFloat(options["margin"], 64)

	if nil == err {
		multiTextOptions.SetMargin(margin)
	}

	// Получаем ширину строки в символах (наверно)
	lineWidth, err := strconv.Atoi(options["lineWidth"])

	if nil == err {
		multiTextOptions.SetLineWidth(lineWidth)
	}

	return multiTextOptions
}
