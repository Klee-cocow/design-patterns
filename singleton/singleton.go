package singleton

import "sync"

type Singleton struct {
	Data string
}

var instance *Singleton
var once sync.Once
var mu sync.Mutex

// GetInstance 普通单例
func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{
			Data: "普通单例数据",
		}
	})
	return instance
}

// GetLazySingleton 懒汉单例
func GetLazySingleton() *Singleton {
	mu.Lock()
	defer mu.Unlock()
	if instance == nil {
		instance = &Singleton{
			Data: "懒汉单例数据",
		}
	}
	return instance
}
