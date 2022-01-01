package enum

import (
	"github.com/mo3golom/wonder-animator/internal/service/processor/processorType"
	"github.com/mo3golom/wonder-animator/pkg/draw2dExtend"
)

type ProcessorType struct {
	Id        string
	Name      string
	Processor processorType.ProcessorInterface
}

func GetProcessorTypes(builder *draw2dExtend.DrawBuilder) []ProcessorType {
	// Абстракный процессор, который содержит все что нужно для работы остальных процессоров
	mainProcessor := &processorType.ProcessorStruct{}
	mainProcessor.DrawBuilder = builder

	textProcessor := &processorType.TextProcessor{ProcessorStruct: mainProcessor}

	return []ProcessorType{
		{Id: "text", Name: "Текст", Processor: textProcessor},
		{Id: "timer", Name: "Таймер", Processor: &processorType.TimerProcessor{TextProcessor: textProcessor}},
		{Id: "rectangle", Name: "Прямоугольник", Processor: &processorType.RectangleProcessor{ProcessorStruct: mainProcessor}},
		{Id: "circle", Name: "Круг", Processor: &processorType.CircleProcessor{ProcessorStruct: mainProcessor}},
		{Id: "triangle", Name: "Треугольник", Processor: &processorType.TriangleProcessor{ProcessorStruct: mainProcessor}},
		{Id: "image", Name: "Изображение", Processor: &processorType.ImageProcessor{ProcessorStruct: mainProcessor}},
	}
}
