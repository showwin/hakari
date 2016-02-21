package main

import (
  "fmt"
  "sync"
  "time"
  "strconv"
  "net/http"
	"net/url"
  "io/ioutil"
  "log"

  "gopkg.in/yaml.v2"
)

var scenario = make(map[interface{}]interface{})

func loadScenario() {
  file, err := ioutil.ReadFile("scenario.yaml")
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(file, &scenario)
	if err != nil {
		log.Fatal(err)
	}
  // return m.(map[interface{}]interface{})
	//for key, value := range a["Header"].(map[interface {}]interface {}) {
	//  fmt.Printf("m: %v::%v\n", key, value)
	//}
}

func Scenario(wg *sync.WaitGroup, m *sync.Mutex, finishTime time.Time) {
	score := 0
	resp := 200
	var c []*http.Cookie

  for t, opt := range scenario {
    title := t.(string)
    fmt.Printf("scenario start: %v\n", title)
    method := opt.(map[interface {}]interface {})["method"].(string)
    path := opt.(map[interface {}]interface {})["url"].(string)
    params := opt.(map[interface {}]interface {})["parameter"]
    v := url.Values{}
    if params != nil {
      for key, val := range params.(map[interface {}]interface {}) {
        v.Add(key.(string), val.(string))
      }
    }
    resp, c = HttpRequest(method, path, v, c)
    fmt.Println(title+": "+strconv.Itoa(resp))
    score = CalcScore(score, resp)
  }


  /*
  // GET Request
	resp, c = HttpRequest("GET", "http://54.238.241.177/", nil, c)
  fmt.Println("get index: "+strconv.Itoa(resp))
	score = CalcScore(score, resp)

  // POST Request
  v := url.Values{}
	v.Add("email", "ishocon@isho.con")
  v.Add("password", "ishoconpass")
	resp, c = HttpRequest("POST", "http://54.238.241.177/login", v, c)
  fmt.Println("Login: "+strconv.Itoa(resp))
	score = CalcScore(score, resp)

  // POST Request
	resp, c = HttpRequest("POST", "http://54.238.241.177/products/buy/9999", nil, c)
  fmt.Println("Buy: "+strconv.Itoa(resp))
	score = CalcScore(score, resp)
  */

	UpdateScore(score, wg, m, finishTime)
}
