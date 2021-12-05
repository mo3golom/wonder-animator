package processor

import (
	"github.com/mo3golom/wonder-animator/internal/dto"
	"github.com/mo3golom/wonder-animator/internal/dto/enum"
	"github.com/mo3golom/wonder-animator/internal/service/processor/processorType"
	"image/draw"
)

type Handler struct {
	id          string
	nextHandler *Handler
	processor   processorType.ProcessorInterface
}

func NewProcessorHandlerBus(blockTypes []enum.ProcessorType) *Handler {
	var prevTypeHandler *Handler

	for _, blockType := range blockTypes {
		handler := &Handler{id: blockType.Id(), processor: blockType.Processor()}

		if nil != prevTypeHandler {
			handler.nextHandler = prevTypeHandler
		}

		prevTypeHandler = handler
	}

	return prevTypeHandler
}

func (th *Handler) Handle(dest draw.Image, block *dto.Block, frameData *dto.FrameData) (err error) {
	blockType := block.Type

	if blockType.Id == th.id {
		return th.processor.TransformOptions(&blockType.Options).Processing(dest, block, frameData)
	}

	if nil != th.nextHandler {
		return th.nextHandler.Handle(dest, block, frameData)
	}

	return nil
}
