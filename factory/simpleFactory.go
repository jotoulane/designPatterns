package factory

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
)

type IRuleConfigParser interface {
	Parse(data []byte)
}

type JsonRuleConfigParser struct {
}

func (j JsonRuleConfigParser) Parse(data []byte) {
	var jsonData map[string]interface{}
	err := json.Unmarshal(data, &jsonData)
	if err != nil {
		fmt.Printf("Error decoding JSON:%v\n", err)
		return
	}
	fmt.Printf("jsonData:%v\n", jsonData)
}

type YamlRuleConfigParser struct {
}

func (y YamlRuleConfigParser) Parse(data []byte) {
	var yamlData map[string]interface{}
	err := yaml.Unmarshal(data, &yamlData)
	if err != nil {
		fmt.Println("Error decoding YAML:", err)
		return
	}
	fmt.Printf("yamlData:%v\n", yamlData)
}

func NewIRuleConfigParser(t string) IRuleConfigParser {
	switch t {
	case "json":
		return JsonRuleConfigParser{}
	case "yaml":
		return YamlRuleConfigParser{}
	default:
		return nil
	}
}
