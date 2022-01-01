package processorOptions

type ImageOptions struct {
	Data   string       `mapstructure:"data"`
	Width  int          `mapstructure:"width"`
	Height int          `mapstructure:"height"`
	Mask   *MaskOptions `mapstructure:"mask"`
}

func NewImageOptions() *ImageOptions {
	return &ImageOptions{
		Width:  0,
		Height: 0,
	}
}
