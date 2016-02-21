package main

import (
  //"fmt"
  "sync"
  "time"
  "net/http"
	"net/url"
  "io/ioutil"
  "log"

  "gopkg.in/yaml.v2"
)

type Parameter struct {
  Key string
  Value string
}

type Request struct {
  Title string
  Method string
  Url string
  Params []Parameter
}

type Scenario []Request

var scenario = Scenario{}

func loadScenario() {
  file, err := ioutil.ReadFile("scenario.yaml")
	if err != nil {
		log.Fatal(err)
	}

  m := yaml.MapSlice{}
	err = yaml.Unmarshal(file, &m)
	if err != nil {
		log.Fatal(err)
	}

  for _, r := range m {
    req := Request{}
    req.Title = r.Key.(string)
    for _, o := range r.Value.(yaml.MapSlice) {
      switch o.Key {
      case "method":
        req.Method = o.Value.(string)
      case "url":
        req.Url = o.Value.(string)
      case "parameter":
        for _, p := range o.Value.(yaml.MapSlice) {
          par := Parameter{}
          par.Key = p.Key.(string)
          par.Value = p.Value.(string)
          req.Params = append(req.Params, par)
        }
      }
    }
    scenario = append(scenario, req)
  }
  //fmt.Println(scenario)
}

func StartScenario(wg *sync.WaitGroup, m *sync.Mutex, finishTime time.Time) {
  var status int
  var c []*http.Cookie

  for _, r := range scenario {
    //fmt.Printf("scenario start: %v\n", r.Title)

    v := url.Values{}
    for _, p := range r.Params {
      v.Add(p.Key, p.Value)
    }

    sTime := time.Now()
    status, c = HttpRequest(r.Method, r.Url, v, c)
    fTime := time.Now()
    Record(r.Title, status, fTime.Sub(sTime), m)
    //fmt.Println("\n")
  }
  CheckFinish(wg, finishTime)
}
