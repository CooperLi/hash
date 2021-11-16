package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/CooperLi/hash/pkg/util"
)

const (
	// SHA1      = "sha1"
	MD5       = "md5"
	SHA256    = "sha256"
	SHA256Old = "sha256old"
	SHA256New = "sha256new"
	IMOHash   = "imohash"
)

type option struct {
	path string
	algo string
	gc   bool
	op   int
}

func main() {
	var opt option
	flag.StringVar(&opt.path, "path", "", "")
	flag.StringVar(&opt.algo, "algo", "", "")
	flag.BoolVar(&opt.gc, "gc", false, "")
	flag.IntVar(&opt.op, "op", 1, "")

	flag.Parse()

	if len(flag.Args()) != 0 {
		fmt.Println("accept no args")
		os.Exit(1)
	}
	if opt.op < 1 {
		fmt.Println("operation time can not less than one")
		os.Exit(1)
	}
	algoList := []string{MD5, IMOHash, SHA256, SHA256New, SHA256Old}

	if len(opt.algo) == 0 {
		for _, al := range algoList {
			var hash string
			startTime := time.Now()
			util.TraceMemStats()
			for i := 0; i < opt.op; i++ {
				nowHash, err := Hash(opt.path, al)
				if err != nil {
					fmt.Println(err.Error())
				}
				if opt.gc {
					util.TraceMemStats()
					runtime.GC()
				}
				hash = nowHash
			}
			endTime := time.Now()
			util.TraceMemStats()
			delta := endTime.Sub(startTime).Seconds() / float64(opt.op)
			averageTime, _ := time.ParseDuration(fmt.Sprintf("%fs", delta))
			fmt.Printf("%s: %s: %s\n", al, averageTime.String(), hash)
		}
	} else {
		var hash string
		startTime := time.Now()
		util.TraceMemStats()
		for i := 0; i < opt.op; i++ {
			nowHash, err := Hash(opt.path, opt.algo)
			if err != nil {
				fmt.Println(err.Error())
			}
			if opt.gc {
				util.TraceMemStats()
				runtime.GC()
			}
			hash = nowHash
		}
		endTime := time.Now()
		util.TraceMemStats()
		delta := endTime.Sub(startTime).Seconds() / float64(opt.op)
		averageTime, _ := time.ParseDuration(fmt.Sprintf("%fs", delta))
		fmt.Printf("%s: %s: %s\n", opt.algo, averageTime.String(), hash)
	}
}
