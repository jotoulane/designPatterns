package factory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewIRuleConfigParser(t *testing.T) {

	t.Run("should return JsonRuleConfigParser when type is json", func(t *testing.T) {
		parser := NewIRuleConfigParser("json")
		_, ok := parser.(JsonRuleConfigParser)
		assert.True(t, ok)
	})

	t.Run("should return YamlRuleConfigParser when type is yaml", func(t *testing.T) {
		parser := NewIRuleConfigParser("yaml")
		_, ok := parser.(YamlRuleConfigParser)
		assert.True(t, ok)
	})

	t.Run("should return nil when type is invalid", func(t *testing.T) {
		parser := NewIRuleConfigParser("invalid")
		assert.Nil(t, parser)
	})
}

func TestJson(t *testing.T) {
	parser := NewIRuleConfigParser("json")
	if parser == nil {
		t.Errorf("Expected non-nil value, but got nil")
	}

	data := []byte(`{"foo": "bar"}`)

	parser.Parse(data)
}

func TestYaml(t *testing.T) {
	parser := NewIRuleConfigParser("yaml")
	if parser == nil {
		t.Errorf("Expected non-nil value, but got nil")
	}

	data := []byte(`foo: bar`)

	parser.Parse(data)
}
