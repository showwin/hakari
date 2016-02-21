package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

var httpHeader = make(map[interface{}]interface{})

func LoadConfig(path string) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		ShowLog("WARNING: Cannot open " + path + ". hakari will use default HTTP client")
		return
	}

	m := make(map[interface{}]interface{})
	err = yaml.Unmarshal(file, &m)
	if err != nil {
		log.Fatal(err)
	}

	loadHttpHeader(m)
}

func loadHttpHeader(m map[interface{}]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			log.Fatal("Invalid Format in Config File")
		}
	}()
	httpHeader = m["Header"].(map[interface{}]interface{})
}
