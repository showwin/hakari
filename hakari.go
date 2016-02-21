package main

import (
	"flag"
	"fmt"
	"strconv"
	"time"
  "sync"

	_ "github.com/go-sql-driver/mysql"
)

func LoopRequests(wg *sync.WaitGroup, m *sync.Mutex, finishTime time.Time) {
	//for {
		Scenario(wg, m, finishTime)
	//}
}

func StartStressTest(workload int) {
	httpHeader = LoadHttpHeader()
	ShowLog("Stress Test Start!  Workload: " + strconv.Itoa(workload))
	finishTime := time.Now().Add(1 * time.Minute)
	wg := new(sync.WaitGroup)
	m := new(sync.Mutex)
	for i := 0; i < workload; i++ {
		wg.Add(1)
		go LoopRequests(wg, m, finishTime)
	}
	wg.Wait()
}

func ShowLog(str string) {
	fmt.Println(time.Now().Format("15:04:05") + "  " + str)
}

var httpHeader = make(map[interface{}]interface{})
var TotalScore = 0
var Finished = false

func main() {
	flag.Usage = func() {
		fmt.Println(`Usage: ./hakari [option]
Options:
  --workload N	    Run with N workloads
  --c FILE          Config File`)
	}

	var (
		workload = flag.Int("workload", 1, "run benchmark with n workloads")
	)
	flag.Parse()

	StartStressTest(*workload)
}
