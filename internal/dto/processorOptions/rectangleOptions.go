package processorOptions

type RectangleOptions struct {
	Width         float64 `mapstructure:"width"`
	Height        float64 `mapstructure:"height"`
	RoundedRadius float64 `mapstructure:"roundedRadius"`
	StrokeWidth   float64 `mapstructure:"strokeWidth"`
	FillColor     string  `mapstructure:"fillColor"`
	StrokeColor   string  `mapstructure:"strokeColor"`
}

func NewRectangleOptions() *RectangleOptions {
	return &RectangleOptions{
		Width:         0,
		Height:        0,
		RoundedRadius: 0,
		StrokeWidth:   0,
		FillColor:     "#ffffff",
		StrokeColor:   "#000000",
	}
}
