package factory_method

import "fmt"

// 智能灯统一接口
type ILight interface {
	// 智能灯的基本信息
	ID() int
	Name() string

	// 智能灯的基本操作
	Open() error
	Close() error
}

// 智能灯的生产工厂统一接口
type ILightFactory interface {
	// 根据智能灯的基本配置信息来生产一个智能灯
	Create(info *LightInfo) (error, ILight)
}

// 智能灯的基本配置信息
type LightInfo struct {
	xID     int
	xName   string
	xVendor string
	xModel  string
}

// 智能灯配置实例创建
func newLightInfo(id int, name, vendor, model string) *LightInfo {
	return &LightInfo{
		xID:     id,
		xName:   name,
		xVendor: vendor,
		xModel:  model,
	}
}

// 实现 ILight 接口
func (l *LightInfo) ID() int {
	return l.xID
}

func (l *LightInfo) Name() string {
	return l.xName
}

func (l *LightInfo) Vendor() string {
	return l.xVendor
}

var defaultFactoryRegistry = newFactoryRegistry()

// 维护厂商名称与厂商的智能灯生产厂的关系的接口
type IFactoryRegistry interface {
	// 注册新的厂商
	Set(vendor string, factory ILightFactory)
	// 根据厂商的名称来获取厂商的智能灯生产厂
	Get(vendor string) ILightFactory
}

// 实现 IFactoryRegistry 接口
type simpleFactoryRegistry struct {
	factoryMap map[string]ILightFactory
}

func newFactoryRegistry() IFactoryRegistry {
	return &simpleFactoryRegistry{
		factoryMap: make(map[string]ILightFactory, 0),
	}
}

func (f *simpleFactoryRegistry) Set(vendor string, factory ILightFactory) {
	f.factoryMap[vendor] = factory
}

func (f *simpleFactoryRegistry) Get(vendor string) ILightFactory {
	light, ok := f.factoryMap[vendor]
	if ok {
		return light
	}
	return nil
}

// 不同厂商的智能灯工厂实现
type meHomeLightFactory struct {
}

func newMeHomeLightFactory() ILightFactory {
	return &meHomeLightFactory{}
}

func (me *meHomeLightFactory) Create(info *LightInfo) (error, ILight) {
	return nil, newMeHomeLight(info)
}

type meHomeLight struct {
	LightInfo
}

func newMeHomeLight(info *LightInfo) *meHomeLight {
	return &meHomeLight{
		*info,
	}
}

func (me *meHomeLight) Open() error {
	fmt.Printf("miHomeLight.open, %v\n", &me.LightInfo)
	return nil
}

func (me *meHomeLight) Close() error {
	fmt.Printf("miHomeLight.Close, %v\n", &me.LightInfo)
	return nil
}

type redMeLightFactory struct {
}

func newRedMeLightFactory() ILightFactory {
	return &redMeLightFactory{}
}

func (rm *redMeLightFactory) Create(info *LightInfo) (error, ILight) {
	return nil, newRedMeLight(info)
}

type redMeLight struct {
	LightInfo
}

func newRedMeLight(info *LightInfo) *redMeLight {
	return &redMeLight{
		*info,
	}
}

func (rm *redMeLight) Open() error {
	fmt.Printf("redMeLight.Open, %v\n", &rm.LightInfo)
	return nil
}

func (rm *redMeLight) Close() error {
	fmt.Printf("redMeLight.Close, %v\n", &rm.LightInfo)
	return nil
}
