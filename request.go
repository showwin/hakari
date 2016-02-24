package main

import (
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"time"
)

func hireWorker(jobQueue chan int) {
	var client http.Client
	var ck []*http.Cookie

	for _ = range jobQueue {
		for _, r := range scenario.Requests {
			var status int
			var duration time.Duration

			// set POST parameters
			v := url.Values{}
			for _, p := range r.Params {
				v.Add(p.Key, p.Value)
			}

			status, ck, duration = sendRequest(client, r.Method, r.Url, v, ck)

			Record(r.Title, status, duration)
		}
	}
}

func sendRequest(client http.Client, method string, path string, params url.Values, cookies []*http.Cookie) (int, []*http.Cookie, time.Duration) {
	req, _ := http.NewRequest(method, path, strings.NewReader(params.Encode()))
	for key, value := range config.HttpHeader {
		req.Header.Add(key.(string), value.(string))
	}
	jar, _ := cookiejar.New(nil)
	CookieURL, _ := url.Parse(path)
	jar.SetCookies(CookieURL, cookies)
	client.Jar = jar

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
