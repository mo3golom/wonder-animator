package WonderAnimator

import (
	"github.com/mo3golom/wonder-animator/internal/dto/enum"
	"github.com/mo3golom/wonder-animator/internal/service"
	"github.com/mo3golom/wonder-animator/internal/service/frameSetSaver"
	"github.com/mo3golom/wonder-animator/internal/service/processor"
	"github.com/mo3golom/wonder-animator/pkg/draw2dExtend"
	"github.com/mo3golom/wonder-animator/pkg/loader"
	WonderEffects "github.com/mo3golom/wonder-effects"
	"github.com/mo3golom/wonder-glitch/wonderGlitchDTO"
	"github.com/mo3golom/wonder-glitch/wonderGlitchService"
	"path"
	"runtime"
)

type wonderAnimator struct {
	fontName, fontPath   string
	additionalProcessors []enum.ProcessorType
	additionalEffects    []WonderEffects.EffectType
	additionalGlitches   []wonderGlitchDTO.EffectType
}

func WonderAnimator() *wonderAnimator {
	return &wonderAnimator{}
}

func (wa *wonderAnimator) SetFont(name, path string) *wonderAnimator {
	wa.fontName = name
	wa.fontPath = path

	return wa
}

func (wa *wonderAnimator) addProcessor(processorType enum.ProcessorType) *wonderAnimator {
	wa.additionalProcessors = append(wa.additionalProcessors, processorType)

	return wa
}

func (wa *wonderAnimator) addEffect(effectType WonderEffects.EffectType) *wonderAnimator {
	wa.additionalEffects = append(wa.additionalEffects, effectType)

	return wa
}
func (wa *wonderAnimator) addGlitch(effectType wonderGlitchDTO.EffectType) *wonderAnimator {
	wa.additionalGlitches = append(wa.additionalGlitches, effectType)

	return wa
}

// Generate фасад для облегчения использования пакета, без необходимости вникать как и какие зависимости собиратьыыы
func (wa *wonderAnimator) Generate(inputObject InputObject, saveType string) frameSetSaver.SaverInterface {

	// Если не был установлен шрифт, то пытаемся загрузить шрифт по умолчанию
	if "" == wa.fontName || "" == wa.fontPath {
		_, file, _, ok := runtime.Caller(0)

		if !ok {
			panic("Не удалось загрузить шрифт по умолчанию")
		}

		wa.fontName = "Roboto"
		wa.fontPath = path.Dir(file) + "/resources/Roboto-Regular.ttf"
	}

	fontData := loader.LoadAndRegister(wa.fontName, wa.fontPath)

	// Собираем сервис, который генерирует кадры
	frameCreatorService := service.NewFrameCreatorService(
		inputObject.GetFPS(),
		processor.NewProcessorHandlerBus(
			append(
				enum.GetProcessorTypes(
					draw2dExtend.NewDrawBuilder(*fontData),
				),
				wa.additionalProcessors...,
			),
		),
		WonderEffects.NewEffectHandlerBus(
			append(
				WonderEffects.GetTypesList(),
				wa.additionalEffects...,
			),
		),
		wonderGlitchService.NewGlitchService(
			wonderGlitchService.NewEffectHandlerBus(
				append(
					wonderGlitchDTO.GetTypesList(),
					wa.additionalGlitches...,
				),
			),
		),
	)

	// Собираем фабрику сервисов, которые сохраняют набор кадров
	saverFactory := service.NewFrameSaverFactory(enum.GetFrameSetSaveStrategies())
	// Получаем конкретного "сохранятеля"
	concreteSaver, ok := saverFactory.SaveType(saveType)

	if !ok {
		panic("saver not selected!")
	}

	// Пробуем установить картинку как фон
	backgroundImage := inputObject.BackgroundImage

	if nil != backgroundImage {
		frameCreatorService.WithBackgroundImage(*backgroundImage)
	}

	// Пробуем установить цвет как фон
	backgroundColor := inputObject.BackgroundColor

	if nil != backgroundColor {
		frameCreatorService.WithBackgroundColor(*backgroundColor)
	}

	// Генерируем набор кадров
	frameSet := frameCreatorService.CreateFrameSet(
		inputObject.Blocks,
		inputObject.Width,
		inputObject.Height,
		inputObject.Duration,
	)

	return concreteSaver.SetFramerate(inputObject.GetFPS()).SetFrameSet(&frameSet)
}
