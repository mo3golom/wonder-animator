package processorType

import (
	"github.com/mo3golom/wonder-animator/internal/dto"
	"github.com/mo3golom/wonder-animator/pkg/draw2dExtend"
	WonderEffects "github.com/mo3golom/wonder-effects"
	"github.com/mo3golom/wonder-effects/wonderEffectDTO"
	"image/draw"
)

type ProcessorStruct struct {
	drawBuilder      *draw2dExtend.DrawBuilder
	effectHandlerBus *WonderEffects.Handler
}

func (p *ProcessorStruct) SetDrawBuilder(drawBuilder *draw2dExtend.DrawBuilder) *ProcessorStruct {
	p.drawBuilder = drawBuilder

	return p
}

func (p *ProcessorStruct) SetEffectHandlerBus(effectHandlerBus *WonderEffects.Handler) *ProcessorStruct {
	p.effectHandlerBus = effectHandlerBus

	return p
}

func (p *ProcessorStruct) applyEffects(block *dto.Block, frameData *dto.FrameData) *wonderEffectDTO.EffectValues {
	progress := float32(frameData.Pos) / float32(frameData.Max)
	return p.applyEffectsWidthAdditionalData(block, progress, 0, 0)
}

func (p *ProcessorStruct) applyEffectsWidthAdditionalData(block *dto.Block, progress float32, addX, addY float64) *wonderEffectDTO.EffectValues {
	if nil == block.Effects {
		return nil
	}

	effectValues := wonderEffectDTO.NewEffectValues()
	effectValues.StartX = block.Position.X + addX
	effectValues.StartY = block.Position.Y + addY
	effectValues.StartRotate = block.Rotate
	effectValues.StartOpacity = block.Opacity
	effectValues.StartScale = block.Scale
	effectValues.RotatePoint = block.RotatePoint

	for _, effect := range *block.Effects {
		_ = p.effectHandlerBus.Handle(&effect, effectValues, &progress)
	}

	return effectValues
}

type ProcessorInterface interface {
	Processing(dest draw.Image, block *dto.Block, frameData *dto.FrameData) (err error)
	TransformOptions(options *map[string]string) ProcessorInterface
}
