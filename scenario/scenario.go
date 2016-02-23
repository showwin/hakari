package scenario

import (
	"io/ioutil"
	"log"

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

type Scenario struct {
	Requests []Request
}

func (s *Scenario) Read(path string) {
	defer func() {
		if err := recover(); err != nil {
			log.Fatal("Invalid Format in Scenario File")
		}
	}()

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
		s.Requests = append(s.Requests, req)
	}
}
