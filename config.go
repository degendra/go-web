package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

//Config is a singleton configuration holder
var Config Configuration

//Configuration holds config data
type Configuration struct {
	Level int64
}

//LoadConfig parses file to Configuration struct
func LoadConfig() {
	source, err := ioutil.ReadFile("config.yml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(source, &Config)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Value: %#v\n", Config.Level)
}
