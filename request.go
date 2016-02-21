package main

import (
  "net/http"
	"net/http/cookiejar"
	"net/url"
  "strings"
  "log"
)

func HttpRequest(method string, path string, params url.Values, cookies []*http.Cookie) (int, []*http.Cookie) {
	req, _ := http.NewRequest(method, path, strings.NewReader(params.Encode()))
	for key, value := range httpHeader {
		req.Header.Add(key.(string), value.(string))
	}
	jar, _ := cookiejar.New(nil)
	CookieURL, _ := url.Parse(path)
	jar.SetCookies(CookieURL, cookies)
	client := http.Client{Jar: jar}

	resp, err := client.Do(req)
	if err != nil {
    log.Fatal(err)
		return 500, cookies
	}
	defer resp.Body.Close()

	return resp.StatusCode, jar.Cookies(CookieURL)
}
