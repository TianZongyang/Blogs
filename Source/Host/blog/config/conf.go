package config

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

//config.yaml是总配置,profile可以选择子配置(或者不同模式,子配置项中和总配置冲突的地方以子配置为准)
func InitConf() Config {
	conf := defaultConfig()
	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatalln(err)
		return conf
	}
	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		log.Fatalln(err)
		return conf
	}
	if conf.Profile != "" {
		yamlFile, err := ioutil.ReadFile("config-" + conf.Profile + ".yaml")
		if err != nil {
			log.Fatalln(err)
			return conf
		}
		if err != nil {
			log.Fatalln(err)
			return conf
		}
		err = yaml.Unmarshal(yamlFile, &conf)
		if err != nil {
			log.Fatalln(err)
			return conf
		}
	}
	log.Printf("current env is %s\n", conf.Env)
	return conf
}

type (
	Config struct {
		Profile    string `yaml:"profile"`
		Env        string `yaml:"env"`
		SystemName string `yaml:"name"`
		ServerHost string `yaml:"server-host"`
		DB         struct {
			Dialect            string `yaml:"dialect"`
			Conn               string `yaml:"conn"`
			MaxOpenConnections int    `yaml:"max-open-connections"`
			MaxIdleConnections int    `yaml:"max-idle-connections"`
		} `yaml:"db"`
		Redis struct {
			Host      string `yaml:"host"`
			DB        int    `yaml:"db"`
			Password  string `yaml:"password"`
			Namespace string `yaml:"namespace"`
		} `yaml:"redis"`
		ElasticSearch struct {
			Host     string `yaml:"host"`
			Password string `yaml:"password"`
		} `yaml:"elasticsearch"`
		Log struct {
			ErrorPath  string `yaml:"error-path"`
			WarnPath   string `yaml:"warn-path"`
			InfoPath   string `yaml:"info-path"`
			Level      string `yaml:"level"`
			SplitLevel string `yaml:"split-level"`
		} `yaml:"log"`
		Email struct {
		} `yaml:"email"`
		SystemErrorReminder []string `yaml:"system-error-reminder"`
	}
)

func defaultConfig() Config {
	return Config{
		Profile:    "",
		Env:        "debug",
		SystemName: "",
		ServerHost: "0.0.0.0:8080",
		DB: struct {
			Dialect            string `yaml:"dialect"`
			Conn               string `yaml:"conn"`
			MaxOpenConnections int    `yaml:"max-open-connections"`
			MaxIdleConnections int    `yaml:"max-idle-connections"`
		}{
			MaxOpenConnections: 100,
			MaxIdleConnections: 10,
		},
		Redis: struct {
			Host      string `yaml:"host"`
			DB        int    `yaml:"db"`
			Password  string `yaml:"password"`
			Namespace string `yaml:"namespace"`
		}{},
		ElasticSearch: struct {
			Host     string `yaml:"host"`
			Password string `yaml:"password"`
		}{},
		Log: struct {
			ErrorPath  string `yaml:"error-path"`
			WarnPath   string `yaml:"warn-path"`
			InfoPath   string `yaml:"info-path"`
			Level      string `yaml:"level"`
			SplitLevel string `yaml:"split-level"`
		}{
			Level:      "info",
			SplitLevel: "D",
		},
		Email:               struct{}{},
		SystemErrorReminder: []string{},
	}
}
