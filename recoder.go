package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type Result struct {
	Duration float64
	Count    int
}

var result = make(map[string]map[int]Result)

func ShowResult() {
	ShowLog("StressTest Finish!")
	fmt.Println("Result:")
	for _, r := range scenario {
		fmt.Println(r.Title)
		m := result[r.Title]
		for st, res := range m {
			status := strconv.Itoa(st)
			req := strconv.Itoa(res.Count) + " req"
			tpr := strconv.FormatFloat(res.Duration*1000/float64(res.Count), 'f', 2, 64) + " ms/req"
			fmt.Println("\t" + status + ": " + req + ", " + tpr + "\n")
		}
	}
}

func Record(title string, status int, t time.Duration, m *sync.Mutex) {
	dur := t.Seconds()

	m.Lock()
	defer m.Unlock()
	old_d := result[title][status].Duration
	old_c := result[title][status].Count
	b := map[int]Result{status: {Duration: old_d + dur, Count: old_c + 1}}
	result[title] = b
}
