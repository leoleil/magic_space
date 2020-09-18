package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Mysql struct {
	User   string `yaml:"user"`
	Pwd    string `yaml:"pwd"`
	Host   string `yaml:"host"`
	Port   string `yaml:"port"`
	Dbname string `yaml:"dbname"`
}
type Video struct {
	Path string `yaml:"path"`
}
type Email struct {
	User string `yaml:"user"`
	Pwd  string `yaml:"pwd"`
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}
type Host struct {
	Name string `yaml:"name"`
	Port string `yaml:"port"`
}
type App struct {
	Mysql Mysql `yaml:"mysql"`
	Video Video `yaml:"video"`
	Email Email `yaml:"email"`
	Host  Host  `yaml:"host"`
}

func (c *App) GetConf(path string) *App {
	yamlFile, err := ioutil.ReadFile(path)
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
