package processorOptions

const (
	MultiTextModeAll   = "all"
	MultiTextModeAfter = "after"
)

type MultiTextOptions struct {
	*TextOptions
	mode      string
	margin    float64
	lineWidth int
}

func NewMultiTextOptions(textOptions *TextOptions) *MultiTextOptions {
	return &MultiTextOptions{
		mode:        MultiTextModeAfter,
		lineWidth:   28,
		margin:      0,
		TextOptions: textOptions,
	}
}

func (m *MultiTextOptions) Mode() string {
	return m.mode
}

func (m *MultiTextOptions) SetMode(mode string) {
	m.mode = mode
}

func (m *MultiTextOptions) Margin() float64 {
	return m.margin
}

func (m *MultiTextOptions) SetMargin(margin float64) {
	m.margin = margin
}

func (m *MultiTextOptions) LineWidth() int {
	return m.lineWidth
}

func (m *MultiTextOptions) SetLineWidth(lineWidth int) {
	m.lineWidth = lineWidth
}
