package main

import (
	WonderAnimator "github.com/mo3golom/wonder-animator"
	"log"
	"time"

	"encoding/json"
	"fmt"

	"runtime"
)

func main() {
	start := time.Now()

	var inputObject WonderAnimator.InputObject
	decodeErr := json.Unmarshal(

		[]byte(`{"duration":5,"width":176,"height":176,"FPS":12,"BackgroundColor":"#b399f1","blocks":[{"type":{"id":"text","options":{"text":"test","width":"64","height":"64","fillColor":"#DA4828","strokeColor":"#EDED49","strokeWidth":"3","padding":"10","radius":"10","margin":"4","lineWidth":"28"}},"effects":[{"id":"move","options":{"distance": "100", "direction": "left"}}],"startAt":0,"duration":5,"position":{"x":50,"y":50},"opacity":1,"rotate":0,"scale":1,"rotatePoint":"right-bottom"}]}`),
		&inputObject,
	)

	if decodeErr != nil {
		panic(decodeErr)
	}

	_, err := WonderAnimator.WonderAnimator().Generate(inputObject, "webp").SaveInFile("test.webp")

	PrintMemUsage()
	if nil != err {
		log.Println(err)
	}

	elapsed := time.Since(start)
	log.Printf("exec time %s", elapsed)
}

// PrintMemUsage outputs the current, total and OS memory being used. As well as the number
// of garage collection cycles completed.
func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
