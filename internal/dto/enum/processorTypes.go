package enum

import (
	"github.com/mo3golom/wonder-animator/internal/service/processor/processorType"
	"github.com/mo3golom/wonder-animator/pkg/draw2dExtend"
	WonderEffects "github.com/mo3golom/wonder-effects"
)

type ProcessorType struct {
	id        string
	name      string
	processor processorType.ProcessorInterface
}

func (b *ProcessorType) Id() string {
	return b.id
}

func (b *ProcessorType) Name() string {
	return b.name
}

func (b *ProcessorType) Processor() processorType.ProcessorInterface {
	return b.processor
}

func GetProcessorTypes(builder *draw2dExtend.DrawBuilder, effectHandlerBus *WonderEffects.Handler) []ProcessorType {
	// Абстракный процессор, который содержит все что нужно для работы остальных процессоров
	mainProcessor := &processorType.ProcessorStruct{}
	mainProcessor.
		SetDrawBuilder(builder).
		SetEffectHandlerBus(effectHandlerBus)

	return []ProcessorType{
		{id: "text", name: "Текст", processor: &processorType.TextProcessor{ProcessorStruct: mainProcessor}},
		{id: "multiText", name: "Мульти текст", processor: &processorType.MultiTextProcessor{ProcessorStruct: mainProcessor}},
		{id: "timer", name: "Таймер", processor: &processorType.TimerProcessor{ProcessorStruct: mainProcessor}},
		{id: "rectangle", name: "Прямоугольник", processor: &processorType.RectangleProcessor{ProcessorStruct: mainProcessor}},
		{id: "circle", name: "Круг", processor: &processorType.CircleProcessor{ProcessorStruct: mainProcessor}},
		{id: "triangle", name: "Треугольник", processor: &processorType.TriangleProcessor{ProcessorStruct: mainProcessor}},
		{id: "image", name: "Изображение", processor: &processorType.ImageProcessor{ProcessorStruct: mainProcessor}},
	}
}
