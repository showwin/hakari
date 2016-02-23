package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	HttpHeader map[interface{}]interface{}
}

func (c *Config) Read(path string) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Print("WARNING: Cannot open " + path + ". hakari will use default HTTP client")
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
