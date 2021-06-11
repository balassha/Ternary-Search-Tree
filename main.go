package main

import (
	"flag"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"tst2/tst"

	"github.com/pkg/profile"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")

func main() {
	//CPUProfile()
	defer profile.Start().Stop()

	var tr tst.Trie
	tr.Put("Tex-Mex Tilapia", 0)
	tr.Put("Cajun-Spiced Pulled Pork", 0)
	tr.Put("Honey Sesame Chicken", 0)
	tr.Put("Cajun-Spiced Pulled Pork", 0)
	tr.Put("Cherry Balsamic Pork Chops", 0)
	tr.Put("Cajun-Spiced Pulled Pork", 0)
	tr.Put("Hot Honey Barbecue Chicken Legs", 0)
	tr.Traverse()
	//MemProfile()
}

func CPUProfile() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close()
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}
}

func MemProfile() {
	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer f.Close()
		runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
	}
}
