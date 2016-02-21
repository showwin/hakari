package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func GetRequest(c []*http.Cookie, page int) (int, []*http.Cookie) {
	return HttpRequest("GET", "", nil, c)
}

func PostRequest(c []*http.Cookie, productId int) (int, []*http.Cookie) {
	v := url.Values{}
	v.Add("content", "parameter value")
	return HttpRequest("POST", "/post_endpoint", v, c)
}

func HttpRequest(method string, path string, params url.Values, cookies []*http.Cookie) (int, []*http.Cookie) {
	req, _ := http.NewRequest(method, host+path, strings.NewReader(params.Encode()))
	for key, value := range httpHeader {
		req.Header.Add(key.(string), value.(string))
	}
	jar, _ := cookiejar.New(nil)
	CookieURL, _ := url.Parse(host + path)
	jar.SetCookies(CookieURL, cookies)
	client := http.Client{Jar: jar}

	resp, err := client.Do(req)
	if err != nil {
		return 500, cookies
	}
	defer resp.Body.Close()

	return resp.StatusCode, jar.Cookies(CookieURL)
}

func ShowScore() {
	ShowLog("StressTest Finish!")
	ShowLog("Score: " + strconv.Itoa(TotalScore))
	ShowLog("Waiting for Stopping All Workers ...")
}

var httpHeader = make(map[interface{}]interface{})
var host = "http://google.com"
var TotalScore = 0

func StartStressTest(workload int) {
	httpHeader = LoadHttpHeader()
	ShowLog("Stress Test Start!  Workload: " + strconv.Itoa(workload))
	//finishTime := time.Now().Add(1 * time.Minute)
	//wg := new(sync.WaitGroup)
	//m := new(sync.Mutex)
	//for i := 0; i < workload; i++ {
	//	wg.Add(1)
	//	go LoopRequests(wg, m, finishTime)
	//}
	//wg.Wait()
	var c []*http.Cookie
	code, _ := HttpRequest("GET", "", nil, c)
	fmt.Println(code)
}

func ShowLog(str string) {
	fmt.Println(time.Now().Format("15:04:05") + "  " + str)
}

func main() {
	flag.Usage = func() {
		fmt.Println(`Usage: ./hakari [option]
Options:
  --workload N	    Run with N workloads
  --c FILE          Config File`)
	}

	var (
		workload = flag.Int("workload", 1, "run benchmark with n workloads")
	)
	flag.Parse()

	StartStressTest(*workload)
}
