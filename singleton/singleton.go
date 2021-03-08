package singleton

import (
	"fmt"
	"sync"
)

// 单例接口
type IDemoSingleton interface {
	Hello()
}

// 饿汉式单例

type hungrySingleton struct {
}

func newHungrySingleton() *hungrySingleton {
	return &hungrySingleton{}
}

func (h *hungrySingleton) Hello() {
	fmt.Printf("hungrySingleton.Hello\n")
}

var gHungrySingleton IDemoSingleton = newHungrySingleton()

// 双重检查单例子
type doubleCheckSingleton struct {
}

func newDoubleCheckSingleton() *doubleCheckSingleton {
	return &doubleCheckSingleton{}
}

func (d *doubleCheckSingleton) Hello() {
	fmt.Printf("doubleCheckSingleton.Hello\n")
}

var gDoubleCheckSingleton IDemoSingleton = nil
var gSingletonLock = new(sync.Mutex)

func GetDoubleCheckSingleton() IDemoSingleton {
	if gDoubleCheckSingleton == nil {
		gSingletonLock.Lock()
		if gDoubleCheckSingleton == nil {
			gDoubleCheckSingleton = newDoubleCheckSingleton()
		}
		gSingletonLock.Unlock()
	}
	return gDoubleCheckSingleton
}

// 容器式单例
type IBeanController interface {
	GetBean(string) (bool, interface{})
	SetBean(string, interface{}) error
}

type beanContainer struct {
	beans map[string]interface{}
	mu    sync.RWMutex
}

func newBeanContainer() *beanContainer {
	return &beanContainer{
		beans: make(map[string]interface{}),
	}
}

func (b *beanContainer) GetBean(name string) (bool, interface{}) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	it, ok := b.beans[name]
	if ok {
		return true, it
	}
	return false, nil
}

func (b *beanContainer) SetBean(name string, it interface{}) error {
	b.mu.RLock()
	defer b.mu.RUnlock()

	if _, ok := b.beans[name]; ok {
		return nil
	}
	b.beans[name] = it
	return nil
}

type containedSingleton struct{
}

func (c *containedSingleton) Hello() {
	fmt.Printf("containedSingleton.Hello\n")
}

func newContainedSingleton() IDemoSingleton {
	return &containedSingleton{}
}

var defaultBeanContainer = newBeanContainer()
