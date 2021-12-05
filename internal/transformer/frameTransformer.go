package transformer

func SecondsToFrameCount(seconds float32, framesPerSecond int) int {
	return int(seconds) * framesPerSecond
}

func FrameCountToSeconds(frameCount int, framesPerSecond int) float32 {
	return float32(frameCount / framesPerSecond)
}
