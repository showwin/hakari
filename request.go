package main

import (
  "net/http"
	"net/http/cookiejar"
	"net/url"
  "strings"
  "time"
  "log"
)

func HttpRequest(method string, path string, params url.Values, cookies []*http.Cookie) (int, []*http.Cookie, time.Duration) {
	req, _ := http.NewRequest(method, path, strings.NewReader(params.Encode()))
	for key, value := range httpHeader {
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
