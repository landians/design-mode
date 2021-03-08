package simple_factory

import (
	"errors"
	"fmt"
	"strings"
)

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

// LightFactory 的实现集合
var defaultLightFactory = newLightFactory()

// 实现 ILightFactory 的简单工厂 xLightFactory
type xLightFactory struct {
}

func newLightFactory() ILightFactory {
	return &xLightFactory{}
}

func (lf *xLightFactory) Create(info *LightInfo) (error, ILight) {
	switch strings.ToLower(info.xVendor) {
	case "mihome":
		return nil, newMiHomeLight(info)
	case "redme":
		return nil, newRedMeLight(info)
	default:
		return errors.New(fmt.Sprintf("unsupported vendor: %s", info.xVendor)), nil
	}
}

// 不同厂商的智能灯控制实现
type meHomeLight struct {
	// 继承 LightInfo，接下来只需要实现 Open() 和 Close() 方法即可实现 ILight 接口
	LightInfo
}

func newMiHomeLight(info *LightInfo) *meHomeLight {
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
