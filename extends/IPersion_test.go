package extends

import (
	"fmt"
	"testing"
)

func TestIPerson(t *testing.T) {
	eat := NewPerson("John").Eat()
	fmt.Printf("eat:%v\n", eat)
}
