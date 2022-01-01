package transformer

func SecondsToFrameCount(seconds float64, framesPerSecond int) int {
	return int(seconds) * framesPerSecond
}

func FrameCountToSeconds(frameCount int, framesPerSecond int) float64 {
	return float64(frameCount) / float64(framesPerSecond)
}
