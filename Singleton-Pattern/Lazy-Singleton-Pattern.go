package Singleton_Pattern

import (
	"sync"
	"sync/atomic"
)

type lazySingleton struct { }

var instance *lazySingleton
var lock sync.Mutex

func GetLazyInstance() *lazySingleton {
	if instance == nil { // 不是完全原子的
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			instance = &lazySingleton{}
		}
	}
	return instance
}

// 通过使用sync/atomic这个包，我们可以原子化加载并设置一个标志，以标识是否创建实例
var initialized uint32

func GetLazyInstanceAtom() *lazySingleton {
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}

	lock.Lock()
	defer lock.Unlock()
	if instance == nil {
		instance = &lazySingleton{}
		atomic.StoreUint32(&initialized, 1)
	}
	return instance
}

// 通过 sync.Once 来确保创建实例的方法只执行一次
var initOnce sync.Once

func GetLazyInstanceOnce() *lazySingleton {
	initOnce.Do(func() {
		instance = &lazySingleton{}
	})
	return instance
}