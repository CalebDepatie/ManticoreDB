package main

import (
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"
)

var Conf Config

type Config struct {
	Databases []string `yaml:"databases"` //the names of each of your databases
}

func GetConfig(path string) (Config, error) {
	var conf Config
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return conf, err
	}

	err = yaml.Unmarshal(file, &conf)
	if err != nil {
		return conf, err
	}

	return conf, nil
}

func (conf Config) UpdateConfigDBs(s []string) {
	conf.Databases = s
}

func (conf Config) SaveConfig(path string) error {
	data, err := yaml.Marshal(conf)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path, data, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
