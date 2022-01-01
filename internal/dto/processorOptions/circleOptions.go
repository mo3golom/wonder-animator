package processorOptions

type CircleOptions struct {
	Radius      float64 `mapstructure:"radius"`
	StrokeWidth float64 `mapstructure:"strokeWidth"`
	FillColor   string  `mapstructure:"fillColor"`
	StrokeColor string  `mapstructure:"strokeColor"`
}

func NewCircleOptions() *CircleOptions {
	return &CircleOptions{
		Radius:      0,
		StrokeWidth: 0,
		FillColor:   "#ffffff",
		StrokeColor: "#000000",
	}
}
