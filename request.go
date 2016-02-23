package main

import (
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"time"
	"sync"
)

func StartScenario (wg *sync.WaitGroup, m *sync.Mutex, finishTime time.Time) {
	var c []*http.Cookie

	for _, r := range scenario.Requests {
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


func HttpRequest(method string, path string, params url.Values, cookies []*http.Cookie) (int, []*http.Cookie, time.Duration) {
	req, _ := http.NewRequest(method, path, strings.NewReader(params.Encode()))
	for key, value := range config.HttpHeader {
		req.Header.Add(key.(string), value.(string))
	}
	jar, _ := cookiejar.New(nil)
	CookieURL, _ := url.Parse(path)
	jar.SetCookies(CookieURL, cookies)
	client := http.Client{Jar: jar}

	sTime := time.Now()
	resp, err := client.Do(req)
	fTime := time.Now()
	t := fTime.Sub(sTime)
	if err != nil {
		log.Fatal(err)
		return 500, cookies, t
	}
	defer resp.Body.Close()

	return resp.StatusCode, jar.Cookies(CookieURL), t
}
