package processorOptions

type ImageOptions struct {
	data          string
	width, height int
}

func NewImageOptions() *ImageOptions {
	return &ImageOptions{
		width:  0,
		height: 0,
	}
}

func (i *ImageOptions) Data() string {
	return i.data
}

func (i *ImageOptions) SetData(data string) {
	i.data = data
}

func (i *ImageOptions) Width() int {
	return i.width
}

func (i *ImageOptions) SetWidth(width int) {
	i.width = width
}

func (i *ImageOptions) Height() int {
	return i.height
}

func (i *ImageOptions) SetHeight(height int) {
	i.height = height
}
