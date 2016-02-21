package main

import (
  "time"
  "strconv"
  "sync"
  "fmt"
  "math"
)

type Result struct {
  Duration float64
  Count int
}

func ShowScore() {
	ShowLog("StressTest Finish!")
	ShowLog("Score: " + strconv.Itoa(TotalScore))
	ShowLog("Waiting for Stopping All Workers ...")
}

func Record(title string, status int, t time.Duration) int {
  dur := math.Ceil(t.Seconds() * 100000) / 100.0
  b := map[int]Result{status: {Duration: dur, Count : 1}}
  result[title] = b
  //result[title][status].duration += t
  fmt.Println(result)
  fmt.Println(dur)
  return 1
}

func UpdateScore(score int, wg *sync.WaitGroup, m *sync.Mutex, finishTime time.Time) {
	m.Lock()
	defer m.Unlock()
	TotalScore = TotalScore + score
	if time.Now().After(finishTime) {
		wg.Done()
		if Finished == false {
			Finished = true
			ShowScore()
		}
	}
}
