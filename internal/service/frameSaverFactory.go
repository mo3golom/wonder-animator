package service

import (
	"github.com/mo3golom/wonder-animator/internal/dto"
	"github.com/mo3golom/wonder-animator/internal/service/frameSetSaver"
)

type FrameSaveFactory struct {
	strategies []dto.FrameSetSaveStrategy
}

func NewFrameSaverFactory(saveStrategies []dto.FrameSetSaveStrategy) *FrameSaveFactory {
	return &FrameSaveFactory{strategies: saveStrategies}
}

func (f *FrameSaveFactory) SaveType(saveType string) (saver frameSetSaver.SaverInterface, ok bool) {
	for _, strategy := range f.strategies {
		if saveType != strategy.Id {
			continue
		}

		saver = strategy.Saver
		ok = true
	}

	return saver, ok
}
