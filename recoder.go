package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type Result struct {
	Status   int
	Duration float64
}

type DurCount struct {
	Duration float64
	Count int
}

var result = make(map[string][]Result)

func ShowResult() {
	for _, r := range scenario.Requests {
		m := result[r.Title]
		rm := map[int]DurCount{}
		for _, res := range m {
			c := rm[res.Status].Count
			d := rm[res.Status].Duration
			dc := DurCount{Duration: d+res.Duration, Count: c+1}
			rm[res.Status] = dc
		}

		fmt.Println(r.Title)
		for st, dc := range rm {
			status := strconv.Itoa(st)
			req := strconv.Itoa(dc.Count) + " req"
			tpr := strconv.FormatFloat(dc.Duration*1000/float64(dc.Count), 'f', 2, 64) + " ms/req"
			fmt.Println("\t" + status + ": " + req + ", " + tpr + "\n")
		}
	}
}

func CreateReport(){
	// ToDo
	fmt.Println("Report will be created.")
}

func Record(title string, status int, t time.Duration, m *sync.Mutex) {
	d := t.Seconds()
	r := Result{Status: status, Duration: d}

	m.Lock()
	defer m.Unlock()
	result[title] = append(result[title], r)
}
