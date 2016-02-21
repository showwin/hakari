package main

import (
	"flag"
	"fmt"
	"strconv"
	"time"
  "sync"
)

func LoopRequests(wg *sync.WaitGroup, m *sync.Mutex, finishTime time.Time) {
	for {
		StartScenario(wg, m, finishTime)
	}
}

func StartStressTest(worker int, cPath string, sPath string, duration int) {
	LoadHttpHeader(cPath)
  LoadScenario(sPath)
	ShowLog("Stress Test Start!  Number of Workers: " + strconv.Itoa(worker))
	finishTime := time.Now().Add(time.Duration(duration) * time.Minute)

	wg := new(sync.WaitGroup)
	m := new(sync.Mutex)
	for i := 0; i < worker; i++ {
		wg.Add(1)
		go LoopRequests(wg, m, finishTime)
	}
	wg.Wait()

  ShowResult()
}

func ShowLog(str string) {
	fmt.Println(time.Now().Format("15:04:05") + "  " + str)
}

func main() {
	flag.Usage = func() {
		fmt.Println(`Usage: ./hakari [option]
Options:
  -w N	           Run with N workers.   default: 2
  -c FILE          Config file.          default: ./config.yaml
  -s FILE          Scenario file.        default: ./scenario.yaml
  -m N             Run for N minutes.    default: 1`)
	}

	var (
		worker = flag.Int("w", 2, "Run with N workers")
    cPath = flag.String("c", "config.yaml", "Config file")
    sPath = flag.String("s", "scenario.yaml", "Scenario file")
    duration = flag.Int("m", 1, "Run for N minutes")
	)
	flag.Parse()

	StartStressTest(*worker, *cPath, *sPath, *duration)
}
