package main

import (
  "io/ioutil"
  "log"

  "gopkg.in/yaml.v2"
)

var httpHeader = make(map[interface{}]interface{})

func LoadHttpHeader(path string) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	m := make(map[interface{}]interface{})
	err = yaml.Unmarshal(file, &m)
	if err != nil {
		log.Fatal(err)
	}
  
  httpHeader = m["Header"].(map[interface{}]interface{})
}
