package processor

import (
	"errors"
	"github.com/mo3golom/wonder-animator/internal/dto"
	"github.com/mo3golom/wonder-animator/internal/dto/enum"
	"github.com/mo3golom/wonder-animator/internal/service/processor/processorType"
	"image"
)

type Handler struct {
	id          string
	nextHandler *Handler
	processor   processorType.ProcessorInterface
}

func NewProcessorHandlerBus(processorTypes []enum.ProcessorType) *Handler {
	var handler *Handler

	for _, pType := range processorTypes {
		handler = handler.AddProcessor(pType)
	}

	return handler
}

func (h *Handler) Handle(block *dto.Block, frameData *dto.FrameData) (output *image.RGBA, err error) {
	blockType := block.Type

	if blockType.Id == h.id {
		return h.processor.TransformOptions(&blockType.Options).Processing(block, frameData)
	}

	if nil != h.nextHandler {
		return h.nextHandler.Handle(block, frameData)
	}

	return nil, errors.New("не удалось найти обработчик для блока")
}

func (h *Handler) AddProcessor(processorType enum.ProcessorType) *Handler {
	newHandler := &Handler{id: processorType.Id, processor: processorType.Processor}
	newHandler.nextHandler = h

	return newHandler
}
