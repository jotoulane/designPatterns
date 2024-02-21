package factory

import "fmt"

type IRuleConfigParser interface {
	Parse(data []byte)
}

type JsonRuleConfigParser struct {
}

func (j JsonRuleConfigParser) Parse(data []byte) {
	fmt.Printf("data:%v\n", data)
}

type YamlRuleConfigParser struct {
}

func (y YamlRuleConfigParser) Parse(data []byte) {
	fmt.Printf("data:%v\n", data)
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
