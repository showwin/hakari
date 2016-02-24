package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"time"

	cfg "github.com/showwin/hakari/config"
	scn "github.com/showwin/hakari/scenario"
)

var (
	worker     = flag.Int("w", 2, "")
	cPath      = flag.String("c", "config.yml", "")
	sPath      = flag.String("s", "scenario.yml", "")
	duration   = flag.Int("d", 10, "")
	report_flg = flag.Bool("report", false, "")

	config   = cfg.Config{}
	scenario = scn.Scenario{}
)

var usage = `Usage: hakari [options...]

Options:
	-w N	           Run with N workers concurrently.   default: 2
	-d N             Run for N seconds.    							default: 10
	-c FILE          Config file.          							default: ./config.yml
	-s FILE          Scenario file.        							default: ./scenario.yml

	--report				 Create detail report.
`

func StartHakari(worker int, duration int) {
	log.Print("hakari Start!  Number of Workers: " + strconv.Itoa(worker))
	finishTime := time.Now().Add(time.Duration(duration) * time.Second)
	jobQueue := make(chan int)

	for i := 0; i < worker; i++ {
		go hireWorker(jobQueue)
	}

	for {
		if time.Now().Before(finishTime) {
			jobQueue <- 1
		} else {
			break
		}
	}

	log.Print("hakari Finish!")
}

func main() {
	flag.Usage = func() {
		fmt.Println(usage)
	}

	flag.Parse()

	config.Read(*cPath)
	scenario.Read(*sPath)

	StartHakari(*worker, *duration)

	if *report_flg == true {
		CreateReport()
	} else {
		ShowResult()
	}
}
