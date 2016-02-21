package main

import (
  "fmt"
  "sync"
  "time"
  "strconv"
  "net/http"
	"net/url"
)

func Scenario(wg *sync.WaitGroup, m *sync.Mutex, finishTime time.Time) {
	score := 0
	resp := 200
	var c []*http.Cookie

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

	UpdateScore(score, wg, m, finishTime)
}
