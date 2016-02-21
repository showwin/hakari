# hakari
A Stress Test Tool for Web Application.  
Easy to Install, Easy to Write Scenario.

## Installation

## Usage

```
$ hakari --help
Usage: hakari [option]
Options:
  -w N	           Run with N workers.   default: 2
  -c FILE          Config file.          default: ./config.yml
  -s FILE          Scenario file.        default: ./scenario.yml
  -m N             Run for N minutes.    default: 1
```

### 1. Write Scenario
hakari use YAML file to write request scenario, as follows:

```yaml
# scenario.yml
TopPage:
  method: "GET"
  url: "http://example.com/"
Login:
  method: "POST"
  url: "http://example.com/login"
  parameter:
    email: "user@example.com"
    password: "secret_password"
BuyProduct:
  method: "POST"
  url: "http://example.com/products/buy/1234"
MyPage:
  method: "GET"
  url: "http://example.com/users/5555"
```

`method` and `url` are **required** for each request.

### 2. Run hakari

```bash
$ hakari
2016/02/21 18:12:47  hakari Start!  Number of Workers: 2
2016/02/21 18:13:48  hakari Finish!
TopPage
	200: 125 req, 238.66 ms/req

Login
	200: 125 req, 255.11 ms/req

BuyProduct
  200: 98 req, 247.68 ms/req
	404: 1 req, 97.12 ms/req
  500: 26 req, 143.82 ms/req

MyPage
	200: 124 req, 233.42 ms/req
```

### (option) Customize HTTP Header 
Require `Header` at top level. Write HTTP header fields freely.
```yaml
# config.yml
Header:
  Accept: "*/*"
  Accept-Encoding: "gzip, deflate, sdch"
  Accept-Language: "ja,en-US;q=0.8,en;q=0.6"
  Cache-Control: "max-age=0"
  User-Agent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/44.0.2403.39 Safari/537.36"
```

## LICENSE

[MIT](https://github.com/showwin/hakari/blob/master/LICENSE)
