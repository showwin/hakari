package main

import (
  "time"
  "strconv"
  "sync"
  "strings"
)

func ShowScore() {
	ShowLog("StressTest Finish!")
	ShowLog("Score: " + strconv.Itoa(TotalScore))
	ShowLog("Waiting for Stopping All Workers ...")
}

func CalcScore(score int, response int) int {
	if response == 200 {
		return score + 1
	} else if strings.Contains(strconv.Itoa(response), "4") {
		return score - 20
	} else {
		return score - 50
	}
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
