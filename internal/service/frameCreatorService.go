package service

import (
	"image"
	"image/color"
	"image/draw"
	"log"
	"math"
	"sync"

	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/mo3golom/wonder-animator/internal/dto"
	"github.com/mo3golom/wonder-animator/internal/service/processor"
	"github.com/mo3golom/wonder-animator/internal/transformer"
	"github.com/mo3golom/wonder-animator/pkg/draw2dExtend"
	"github.com/mo3golom/wonder-animator/pkg/imagingExtend"
	"github.com/mo3golom/wonder-animator/pkg/loader"
	WonderEffects "github.com/mo3golom/wonder-effects"
	"github.com/mo3golom/wonder-effects/wonderEffectDTO"
	"github.com/mo3golom/wonder-glitch/wonderGlitchService"
)

type frameSetItem struct {
	pos   int
	frame draw.Image
}

type FrameCreatorService struct {
	processorHandlerBus *processor.Handler
	effectHandlerBus    *WonderEffects.Handler
	glitchService       *wonderGlitchService.GlitchService
	framesPerSecond     int
	backgroundImage     *image.Image
	backgroundColor     *image.Uniform
}

func NewFrameCreatorService(
	framesPerSecond int,
	processorHandlerBus *processor.Handler,
	effectHandlerBus *WonderEffects.Handler,
	glitchService *wonderGlitchService.GlitchService,
) *FrameCreatorService {
	return &FrameCreatorService{
		framesPerSecond:     framesPerSecond,
		processorHandlerBus: processorHandlerBus,
		effectHandlerBus:    effectHandlerBus,
		glitchService:       glitchService,
	}
}

func (fms *FrameCreatorService) WithBackgroundImage(path string) *FrameCreatorService {
	loadImage, err := loader.LoadImage(path)

	if nil != err {
		log.Println(err)

		return fms
	}

	fms.backgroundImage = &loadImage

	return fms
}

func (fms *FrameCreatorService) WithBackgroundColor(color string) *FrameCreatorService {
	fms.backgroundColor = &image.Uniform{C: draw2dExtend.ParseHexColor(color)}

	return fms
}

func (fms *FrameCreatorService) CreateFrameSet(
	blocks *[]dto.Block,
	width int, height int,
	duration float32,
) []draw.Image {
	var backgroundImage image.Image
	var wg sync.WaitGroup
	var mutex sync.Mutex

	// Добавляем фон, если он был определен
	if nil != fms.backgroundImage {
		backgroundImage = *fms.backgroundImage
	}

	backgroundColor := image.NewUniform(color.RGBA{})

	// Добавляем цвет фона, если он был определен
	if nil != fms.backgroundColor {
		backgroundColor = fms.backgroundColor
	}

	durationInFrames := transformer.SecondsToFrameCount(float64(duration), fms.framesPerSecond)
	frameChannel := make(chan *frameSetItem, durationInFrames)

	// Предварительно создаем срез для записи результатов
	frameSet := make([]draw.Image, durationInFrames)

	// Запускаем "параллельно" генерацию кадров
	// В канал будет записана структура с готовым кадром и его "позицией" в массиве
	for framePos := 1; framePos <= durationInFrames; framePos++ {
		wg.Add(1)
		go fms.createFrameGoroutine(width, height, backgroundColor, backgroundImage, *blocks, framePos, durationInFrames, frameChannel, &wg, &mutex)
	}

	// Запускаем наблюдателя, который закроет канал, когда надо
	go func(wg *sync.WaitGroup, channel chan<- *frameSetItem) {
		wg.Wait()
		close(channel)
	}(&wg, frameChannel)

	// Пишем в нужный индекс полученный кадр
	for frameItem := range frameChannel {
		frameSet[frameItem.pos] = frameItem.frame
	}

	return frameSet
}

// Метод для генерации кадра асинхронно
func (fms *FrameCreatorService) createFrameGoroutine(
	width, height int,
	backgroundColor, backgroundImage image.Image,
	blocks []dto.Block,
	framePos int,
	frameMax int,
	channel chan<- *frameSetItem,
	wg *sync.WaitGroup,
	mutex *sync.Mutex,
) {
	frame := image.NewRGBA(image.Rect(0, 0, width, height))
	// Добавляем цвет фона (по-умолчанию - прозрачный)
	draw.Draw(frame, frame.Bounds(), backgroundColor, image.Point{}, draw.Src)

	// Если есть фон, то добавляем его
	if nil != backgroundImage {
		draw.Draw(frame, frame.Bounds(), backgroundImage, backgroundImage.Bounds().Min, draw.Src)
	}

	concreteBlocks := fms.getConcreteFrames(framePos, blocks, fms.framesPerSecond)

	if 0 < len(concreteBlocks) {
		for _, block := range concreteBlocks {
			frameData := &dto.FrameData{
				FPS: fms.framesPerSecond,
				Pos: framePos,
				Max: frameMax,
			}

			// В этом месте из-за многопроцессорности могут возникать неприятные артефакты, поэтому ставим блокировку мьютексом
			// Альтернативыное решение - runtime.GOMAXPROCS(1)
			mutex.Lock()
			frameItem, err := fms.processorHandlerBus.Handle(&block, frameData)
			mutex.Unlock()

			if nil != err {
				log.Println("Ошибка обработки блока:", block.Type.Id, err)

				continue
			}

			// В этом месте из-за многопроцессорности могут возникать неприятные артефакты, поэтому ставим блокировку мьютексом
			// Альтернативыное решение - runtime.GOMAXPROCS(1)
			mutex.Lock()
			frameItem = fms.applyGlitch(frameItem, &block, frameData)
			mutex.Unlock()

			effectValues := fms.applyEffects(&block, frameData)

			frameItemBounds := frameItem.Bounds()
			// Находим точку поворота
			rotatePoint := draw2dExtend.GetRotatePointByType(
				effectValues.RotatePoint,
				effectValues.X(),
				effectValues.Y(),
				float64(frameItemBounds.Dx()),
				float64(frameItemBounds.Dy()),
			)

			// Устанавливаем прозрачность
			output := imagingExtend.Opacity(frameItem, float64(effectValues.Opacity()))
			output = imagingExtend.RotateAround(
				output,
				effectValues.Rotate(),
				math.Abs(effectValues.X()-rotatePoint.X),
				math.Abs(effectValues.Y()-rotatePoint.Y),
			)

			// В этом месте из-за многопроцессорности могут возникать неприятные артефакты, поэтому ставим блокировку мьютексом
			// Альтернативыное решение - runtime.GOMAXPROCS(1)
			mutex.Lock()
			graphicContext := draw2dimg.NewGraphicContext(frame)
			graphicContext.Translate(effectValues.X(), effectValues.Y())
			graphicContext.Translate(rotatePoint.X, rotatePoint.Y)
			graphicContext.Scale(effectValues.Scale(), effectValues.Scale())
			graphicContext.Translate(-rotatePoint.X, -rotatePoint.Y)
			graphicContext.DrawImage(output)
			mutex.Unlock()
		}

	}

	channel <- &frameSetItem{pos: framePos - 1, frame: frame}
	wg.Done()
}

func (fms *FrameCreatorService) getConcreteFrames(framePos int, blocks []dto.Block, framesPerSecond int) (concreteBlocks []dto.Block) {
	for _, block := range blocks {
		durationInFrames := transformer.SecondsToFrameCount(float64(block.Duration), framesPerSecond)
		startAtInFrames := transformer.SecondsToFrameCount(float64(block.StartAt), framesPerSecond)

		// Если текущая позиция кадра меньше чем старт или больше чем конец, то пропускаем
		if startAtInFrames > framePos || framePos > (startAtInFrames+durationInFrames) {
			continue
		}

		concreteBlocks = append(concreteBlocks, block)
	}

	return
}

func (fms *FrameCreatorService) applyEffects(block *dto.Block, frameData *dto.FrameData) *wonderEffectDTO.EffectValues {
	progress := float32(frameData.Pos) / float32(frameData.Max)

	effectValues := wonderEffectDTO.NewEffectValues()
	effectValues.StartX = block.Position.X
	effectValues.StartY = block.Position.Y
	effectValues.StartRotate = block.Rotate
	effectValues.StartOpacity = block.Opacity
	effectValues.StartScale = block.Scale
	effectValues.RotatePoint = block.RotatePoint

	for _, effect := range block.Effects {
		_ = fms.effectHandlerBus.Handle(&effect, effectValues, &progress)
	}

	return effectValues
}

func (fms *FrameCreatorService) applyGlitch(dest *image.RGBA, block *dto.Block, frameData *dto.FrameData) *image.RGBA {
	progress := float64(frameData.Pos) / float64(frameData.Max)

	return fms.glitchService.SetDest(dest).SetFactor(block.GlitchFactor * progress).Glitchify(block.Glitches)
}
