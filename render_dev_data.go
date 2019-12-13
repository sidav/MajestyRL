package main

import (
	"fmt"
	cw "github.com/sidav/golibrl/console"
	"runtime"
)

var memStrings []string 

func updateMemUsage() {
    var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	memStrings = []string{
        fmt.Sprintf("Alloc      = %v MiB", bToMb(m.Alloc)),
        fmt.Sprintf("HeapAlloc  = %v MiB", bToMb(m.HeapAlloc)),
		fmt.Sprintf("TotalAlloc = %v MiB", bToMb(m.TotalAlloc)),
		fmt.Sprintf("Sys        = %v MiB", bToMb(m.Sys)),
		fmt.Sprintf("NumGC      = %v\n", m.NumGC),
        fmt.Sprintf("Lookups    = %d", m.Lookups),
        fmt.Sprintf("Mallocs    = %d", m.Mallocs),
        fmt.Sprintf("Frees      = %d", m.Frees),
	}
}

func PrintMemUsage() {
    if CURRENT_TICK % 100 == 0 {
        updateMemUsage()
    }
	for i := 0; i < len(memStrings); i++ {
		cw.PutString(memStrings[i], 0, i)
	}

}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
