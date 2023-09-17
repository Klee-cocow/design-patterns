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

// GetDoubleCheckInstance 双重锁单例
func GetDoubleCheckInstance() *Singleton {
	if instance == nil {
		mu.Lock()
		defer mu.Unlock()
		if instance == nil {
			instance = &Singleton{
				Data: "双重锁单例数据",
			}
		}
	}
	return instance
}

// 饿汉
var eInstance = &Singleton{
	Data: "饿汉单例数据",
}

func GetEagerInstance() *Singleton {
	return eInstance
}
