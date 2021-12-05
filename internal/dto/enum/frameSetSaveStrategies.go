package enum

import (
	"github.com/mo3golom/wonder-animator/internal/dto"
	"github.com/mo3golom/wonder-animator/internal/service/frameSetSaver"
	"github.com/mo3golom/wonder-animator/pkg/ffmpeg"
)

func GetFrameSetSaveStrategies() []dto.FrameSetSaveStrategy {
	return []dto.FrameSetSaveStrategy{
		{Id: "video", Saver: frameSetSaver.NewVideoSaverStrategy(ffmpeg.NewFFMpeg())},
		{Id: "webp", Saver: frameSetSaver.NewWebpSaver()},
	}
}
