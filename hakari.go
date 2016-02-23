package main

import (
	"flag"
	"fmt"
	"strconv"
	"sync"
	"time"
	"log"

	cfg "github.com/showwin/hakari/config"
	scn "github.com/showwin/hakari/scenario"
)

var (
	worker   = flag.Int("w", 2, "")
	cPath    = flag.String("c", "config.yml", "")
	sPath    = flag.String("s", "scenario.yml", "")
	duration = flag.Int("d", 60, "")
	report_flg = flag.Bool("report", false, "")

	config = cfg.Config{}
	scenario = scn.Scenario{}
)

var usage = `Usage: hakari [options...]

Options:
	-w N	           Run with N workers.   default: 2
	-c FILE          Config file.          default: ./config.yml
	-s FILE          Scenario file.        default: ./scenario.yml
	-d N             Run for N seconds.    default: 60

	--report				 Create detail report.
`

func StartStressTest(worker int, duration int) {
	log.Print("hakari Start!  Number of Workers: " + strconv.Itoa(worker))
	finishTime := time.Now().Add(time.Duration(duration) * time.Second)

	wg := new(sync.WaitGroup)
	m := new(sync.Mutex)
	wg.Add(worker)
	for i := 0; i < worker; i++ {
		go func() {
			for {
				StartScenario(wg, m, finishTime)
			}
		}()
	}
	wg.Wait()
	log.Print("hakari Finish!")
}

func main() {
	flag.Usage = func() {
		fmt.Println(usage)
	}

	flag.Parse()

	config.Read(*cPath)
	scenario.Read(*sPath)

	StartStressTest(*worker, *duration)

	if *report_flg == true {
		CreateReport()
	} else {
		ShowResult()
	}
}
