package loadYaml

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type yamlStruct struct {
	Url     string   `yaml:"url"`
	Method  string   `yaml:"method"`
	Headers []string `yaml:"headers"`
	Body    string   `yaml:body"`
}

func ReqYamlFile(fileName string) *yamlStruct {

	yamlFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Error reading YAML file: %s\n", err)
		os.Exit(1)
	}

	var yamlParse yamlStruct
	err = yaml.Unmarshal(yamlFile, &yamlParse)
	if err != nil {
		fmt.Printf("Error parsing YAML file: %s\n", err)
		os.Exit(1)
	}

	return &yamlParse
}
