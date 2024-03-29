package singleton

import (
	"fmt"
	"sync"
)

type LazySingleton struct{}

var (
	lazySingleton *LazySingleton
	once          = &sync.Once{}
)

// GetLazyInstance 懒汉式
func GetLazyInstance() *LazySingleton {
	fmt.Printf("lazySingleton:%v\n", lazySingleton)
	if lazySingleton == nil {
		once.Do(func() {
			lazySingleton = &LazySingleton{}
		})
	}
	return lazySingleton
}
