package service

import (
	"github.com/mo3golom/wonder-animator/internal/dto"
	"github.com/mo3golom/wonder-animator/internal/service/processor"
	"github.com/mo3golom/wonder-animator/internal/transformer"
	"github.com/mo3golom/wonder-animator/pkg/draw2dExtend"
	"github.com/mo3golom/wonder-animator/pkg/loader"
	"image"
	"image/color"
	"image/draw"
	"log"
	"sync"
)

type frameSetItem struct {
	pos   int
	frame draw.Image
}

type FrameCreatorService struct {
	processorHandlerBus *processor.Handler
	framesPerSecond     int
	backgroundImage     *image.Image
	backgroundColor     *image.Uniform
}

func NewFrameCreatorService(framesPerSecond int, processorHandlerBus *processor.Handler) *FrameCreatorService {
	return &FrameCreatorService{
		framesPerSecond:     framesPerSecond,
		processorHandlerBus: processorHandlerBus,
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

	durationInFrames := transformer.SecondsToFrameCount(duration, fms.framesPerSecond)
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
			err := fms.processorHandlerBus.Handle(frame, &block, frameData)
			mutex.Unlock()

			if nil == err {
				continue
			}

			log.Println("Ошибка обработки блока", block.Type.Id, err)
		}

	}

	channel <- &frameSetItem{pos: framePos - 1, frame: frame}
	wg.Done()
}

func (fms *FrameCreatorService) getConcreteFrames(framePos int, blocks []dto.Block, framesPerSecond int) (concreteBlocks []dto.Block) {
	for _, block := range blocks {
		durationInFrames := transformer.SecondsToFrameCount(block.Duration, framesPerSecond)
		startAtInFrames := transformer.SecondsToFrameCount(block.StartAt, framesPerSecond)

		// Если текущая позиция кадра меньше чем старт или больше чем конец, то пропускаем
		if startAtInFrames > framePos || framePos > (startAtInFrames+durationInFrames) {
			continue
		}

		concreteBlocks = append(concreteBlocks, block)
	}

	return
}
