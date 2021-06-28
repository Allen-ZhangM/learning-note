package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
)

func main() {
	cpuf, err := os.Create("GoPractice/pprof/cpu_profile")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(cpuf)
	defer pprof.StopCPUProfile()

	j := 0
	for i := 0; i < 10000000000; i++ {
		j++
	}
	fmt.Println(j)
}
