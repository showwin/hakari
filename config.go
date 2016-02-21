package main

import (
  "io/ioutil"
  "log"

  "gopkg.in/yaml.v2"
)

func LoadHttpHeader() (map[interface{}]interface{}) {
	file, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	m := make(map[interface{}]interface{})
	err = yaml.Unmarshal(file, &m)
	if err != nil {
		log.Fatal(err)
	}
  return m["Header"].(map[interface{}]interface{})
}
