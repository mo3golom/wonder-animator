package processorOptions

type TextOptions struct {
	Text            string  `mapstructure:"text"`
	Width           int     `mapstructure:"width"`
	Height          int     `mapstructure:"height"`
	FontSize        float64 `mapstructure:"fontSize"`
	Padding         float64 `mapstructure:"padding"`
	Radius          float64 `mapstructure:"radius"`
	TextColor       string  `mapstructure:"textColor"`
	BackgroundColor string  `mapstructure:"backgroundColor"`
}

func NewTextOptions() *TextOptions {
	return &TextOptions{
		Text:            "",
		FontSize:        12,
		Padding:         0,
		Radius:          0,
		TextColor:       "#ffffff",
		BackgroundColor: "#000000",
	}
}
