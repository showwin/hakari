package main

import (
  "fmt"
  "sync"
  "time"
  "net/http"
	"net/url"
  "io/ioutil"
  "log"

  "gopkg.in/yaml.v2"
)

var scenario = yaml.MapSlice{}

func loadScenario() {
  file, err := ioutil.ReadFile("scenario.yaml")
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(file, &scenario)
	if err != nil {
		log.Fatal(err)
	}
}

func Scenario(wg *sync.WaitGroup, m *sync.Mutex, finishTime time.Time) {
	score := 0
  status := 200
	var c []*http.Cookie
  v := url.Values{}
  method := ""
  path := ""

  for _, r := range scenario {
    title := r.Key.(string)
    fmt.Printf("scenario start: %v\n", title)
    for _, o := range r.Value.(yaml.MapSlice) {
      switch o.Key {
      case "method":
        method = o.Value.(string)
      case "url":
        path = o.Value.(string)
      case "parameter":
        for _, p := range o.Value.(yaml.MapSlice) {
          v.Add(p.Key.(string), p.Value.(string))
        }
      }
    }
    sTime := time.Now()
    status, c = HttpRequest(method, path, v, c)
    fTime := time.Now()
    score = Record(title, status, fTime.Sub(sTime))
    fmt.Println("\n")
  }
	UpdateScore(score, wg, m, finishTime)
}
