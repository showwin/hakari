package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sync"
	"time"

	"gopkg.in/yaml.v2"
)

type Parameter struct {
	Key   string
	Value string
}

type Request struct {
	Title  string
	Method string
	Url    string
	Params []Parameter
}

type Scenario []Request

var scenario = Scenario{}

func LoadScenario(path string) {
	file, err := ioutil.ReadFile(path)
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
}

func StartScenario(wg *sync.WaitGroup, m *sync.Mutex, finishTime time.Time) {
	var c []*http.Cookie

	for _, r := range scenario {
		var status int
		var t time.Duration

		v := url.Values{}
		for _, p := range r.Params {
			v.Add(p.Key, p.Value)
		}

		status, c, t = HttpRequest(r.Method, r.Url, v, c)

		Record(r.Title, status, t, m)
	}
	CheckFinish(wg, finishTime)
}

func CheckFinish(wg *sync.WaitGroup, finishTime time.Time) {
	if time.Now().After(finishTime) {
		wg.Done()
	}
}
