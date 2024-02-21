package extends

import (
	"fmt"
	"testing"
)

func TestExtend(t *testing.T) {
	animal := Animal{Name: "dog"}
	eat := animal.Eat()
	fmt.Printf("eat:%v\n", eat)

	cat := Cat{Animal: &animal}
	eat = cat.Eat()
	fmt.Printf("eat:%v\n", eat)
}
