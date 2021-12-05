package dto

import "github.com/mo3golom/wonder-animator/internal/service/frameSetSaver"

type FrameSetSaveStrategy struct {
	Id    string
	Saver frameSetSaver.SaverInterface
}
