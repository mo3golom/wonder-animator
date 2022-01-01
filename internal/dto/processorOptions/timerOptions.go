package processorOptions

const (
	TimerModeNormal  = "normal"
	TimerModeReverse = "reverse"
)

type TimerOptions struct {
	Mode        string `mapstructure:"mode"`
	IncludeZero bool   `mapstructure:"includeZero"`
	Format      string `mapstructure:"format"`
}

func NewTimerOptions() *TimerOptions {
	return &TimerOptions{
		Mode:        TimerModeNormal,
		IncludeZero: false,
	}
}
