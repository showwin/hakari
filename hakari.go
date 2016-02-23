package main

import (
	"flag"
	"fmt"
	"strconv"
	"sync"
	"time"
	"log"

	"./config"
	"./scenario"
)

func LoopRequests(wg *sync.WaitGroup, m *sync.Mutex, finishTime time.Time) {
	for {
		StartScenario(wg, m, finishTime)
	}
}

func StartStressTest(worker int, duration int) {
	log.Print("hakari Start!  Number of Workers: " + strconv.Itoa(worker))
	finishTime := time.Now().Add(time.Duration(duration) * time.Second)

	wg := new(sync.WaitGroup)
	m := new(sync.Mutex)
	for i := 0; i < worker; i++ {
		wg.Add(1)
		go LoopRequests(wg, m, finishTime)
	}
	wg.Wait()
	log.Print("hakari Finish!")
}

var conf = config.Config{}
var scene = scenario.Scenario{}

func main() {
	flag.Usage = func() {
		fmt.Println(`Usage: hakari [option]
Options:
  -w N	           Run with N workers.   default: 2
  -c FILE          Config file.          default: ./config.yml
  -s FILE          Scenario file.        default: ./scenario.yml
  -d N             Run for N seconds.    default: 60
	--report				 Create detail report.`)
	}

	var (
		worker   = flag.Int("w", 2, "Run with N workers")
		cPath    = flag.String("c", "config.yml", "Config file")
		sPath    = flag.String("s", "scenario.yml", "Scenario file")
		duration = flag.Int("d", 60, "Run for N seconds")
		report_flg = flag.Bool("report", false, "Create detail report")
	)
	flag.Parse()

	conf.Read(*cPath)
	scene.Read(*sPath)

	StartStressTest(*worker, *duration)
	if *report_flg == true {
		CreateReport()
	} else {
		ShowResult()
	}
}
