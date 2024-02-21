package singleton

import (
	"fmt"
	"sync"
	"testing"
)

func TestLazySingleton(t *testing.T) {
	lazyInstance1 := GetLazyInstance()
	lazyInstance2 := GetLazyInstance()
	fmt.Printf("instance1: %p\n", lazyInstance1)
	fmt.Printf("instance2: %p\n", lazyInstance2)
	fmt.Printf("is same: %v\n", lazyInstance1 == lazyInstance2)
}

func TestOnce(t *testing.T) {
	var once sync.Once
	once.Do(func() {
		fmt.Println("Only once")
	})
	once.Do(func() {
		fmt.Println("Only once")
	})
}
