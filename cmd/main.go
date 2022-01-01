package main

import (
	"log"
	"time"

	"github.com/mitchellh/mapstructure"
	WonderAnimator "github.com/mo3golom/wonder-animator"

	"encoding/json"
	"fmt"

	"runtime"
)

func main() {
	start := time.Now()

	var inputObjectMap map[string]interface{}
	var inputObject WonderAnimator.InputObject
	decodeErr := json.Unmarshal(
		[]byte(`{"duration":5,"width":250,"height":249,"FPS":60,"BackgroundColor":"#b399f1","blocks":[{"processor":{"id":"image","options":{"text":"ехал грека","mask":{"mask":"mask.png","src":"unnamed.jpg","x":50,"y":50},"data":"vinyl.jpg","format":"15:04:05.00","width":250,"height":249,"fillColor":"#DA4828","strokeColor":"#EDED49","strokeWidth":3,"padding":10,"radius":10,"margin":4,"lineWidth":28}},"startAt":0,"duration":5,"position":{"x":0,"y":0},"opacity":1,"rotate":0,"scale":1,"rotatePoint":"right-bottom"}]}`),
		&inputObjectMap,
	)

	if decodeErr != nil {
		panic(decodeErr)
	}

	decodeErr = mapstructure.Decode(inputObjectMap, &inputObject)

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
