package singleton

import (
	"fmt"
	"testing"
)

func TestSingleton(t *testing.T) {
	instance1 := GetInstance()
	instance2 := GetInstance()
	fmt.Printf("instance1: %p\n", instance1)
	fmt.Printf("instance2: %p\n", instance2)
	fmt.Printf("is same: %v\n", instance1 == instance2)

}
