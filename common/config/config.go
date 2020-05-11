package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)
type Mysql struct {
	User string `yaml:"user"`
	Pwd string `yaml:"pwd"`
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	Dbname string `yaml:"dbname"`
}
type App struct {
	Mysql Mysql `yaml:"mysql"`
}

func (c *App) GetConf() *App {
	yamlFile, err := ioutil.ReadFile("config/app.yml")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println(err.Error())
	}
	return c
}
var AppHandle = &App{}

