package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	HttpHeader map[interface{}]interface{}
	Path string
}

func (c *Config) Read() {
	file, err := ioutil.ReadFile(c.Path)
	if err != nil {
		ShowLog("WARNING: Cannot open " + c.Path + ". hakari will use default HTTP client")
		return
	}

	m := make(map[interface{}]interface{})
	err = yaml.Unmarshal(file, &m)
	if err != nil {
		log.Fatal(err)
	}

	c.HttpHeader = readHttpHeader(m)
}

func readHttpHeader(m map[interface{}]interface{}) (map[interface{}]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			log.Fatal("Invalid Format in Config File")
		}
	}()
	return m["Header"].(map[interface{}]interface{})
}
