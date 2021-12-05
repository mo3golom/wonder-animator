package processorOptions

const (
	TimerModeNormal int = iota
	TimerModeReverse
)

type TimerOptions struct {
	*TextOptions
	mode        int
	includeZero bool
}

func NewTimerOptions(textOptions *TextOptions) *TimerOptions {
	return &TimerOptions{
		mode:        TimerModeNormal,
		includeZero: false,
		TextOptions: textOptions,
	}
}

func (t *TimerOptions) Mode() int {
	return t.mode
}

func (t *TimerOptions) SetMode(mode int) {
	t.mode = mode
}

func (t *TimerOptions) IncludeZero() bool {
	return t.includeZero
}

func (t *TimerOptions) SetIncludeZero(includeZero bool) {
	t.includeZero = includeZero
}
