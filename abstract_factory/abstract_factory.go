package abstract_factory

import "fmt"

// 抽象智能设备的基础接口, 意思是所有的智能设备都有id, name，以及可以进行 open, close 操作
type ISmartDevice interface {
	ID() int
	Name() string

	Open() error
	Close() error
}

// 抽象智能灯的接口，继承了 ISmartDevice 接口，增加灯光模式的控制方法
type ILight interface {
	ISmartDevice

	GetLightMode() (error, int)
	SetLightMode(mode int) error
}

// 抽象空调的接口，继承了 ISmartDevice 接口，增加了温度的控制方法
type IAirConditioner interface {
	ISmartDevice

	GetTemperature() (error, float64)
	SetTemperature(t float64) error
}

// 智能设备工厂接口
type ISmartFactory interface {
	CreateLight(info *DeviceInfo) (error, ILight)
	CreateAirConditioner(info *DeviceInfo) (error, IAirConditioner)
}

// 智能设备的基本信息
type DeviceType string

const DeviceTypeLight DeviceType = "Light"
const DeviceTypeAirConditioner DeviceType = "AirConditioner"

type DeviceInfo struct {
	xID     int
	xName   string
	xType   DeviceType
	xVendor string
	xModel  string
}

func NewDeviceInfo(id int, name string, deviceType DeviceType, vendor string, model string) *DeviceInfo {
	return &DeviceInfo{
		xID:     id,
		xName:   name,
		xType:   deviceType,
		xVendor: vendor,
		xModel:  model,
	}
}

func (d *DeviceInfo) ID() int {
	return d.xID
}

func (d *DeviceInfo) Name() string {
	return d.xName
}

func (d *DeviceInfo) Vendor() string {
	return d.xVendor
}

func (d *DeviceInfo) DeviceType() DeviceType {
	return d.xType
}

var defaultFactoryRegistry = newFactoryRegistry()

// 厂商名称对应厂商的智能灯工厂的注册表
type IFactoryRegistry interface {
	Set(vendor string, factory ISmartFactory)
	Get(vendor string) ISmartFactory
}

type simpleFactorRegistry struct {
	factoryMap map[string]ISmartFactory
}

func newFactoryRegistry() IFactoryRegistry {
	return &simpleFactorRegistry{
		factoryMap: make(map[string]ISmartFactory, 0),
	}
}

func (f *simpleFactorRegistry) Set(vendor string, factory ISmartFactory) {
	f.factoryMap[vendor] = factory
}

func (f *simpleFactorRegistry) Get(vendor string) ISmartFactory {
	smartDevice, ok := f.factoryMap[vendor]
	if ok {
		return smartDevice
	}
	return nil
}

// 不同厂商的工厂，用于实现 SmartFactory 接口
type meHomeFactory struct {
}

func newMeHomeFactory() ISmartFactory {
	return &meHomeFactory{}
}

func (me *meHomeFactory) CreateLight(info *DeviceInfo) (error, ILight) {
	return nil, newMeHomeLight(info)
}

func (me *meHomeFactory) CreateAirConditioner(info *DeviceInfo) (error, IAirConditioner) {
	return nil, newMeHomeAirConditioner(info)
}

// meHome 实现的生产智能灯和智能冰箱
type meHomeLight struct {
	// 基本的设备信息
	DeviceInfo
	// 灯光模式
	xMode int
}

func newMeHomeLight(info *DeviceInfo) *meHomeLight {
	return &meHomeLight{
		DeviceInfo: *info,
		xMode:      0,
	}
}

func (me *meHomeLight) Open() error {
	fmt.Printf("miHomeLight.open, %v\n", &me.DeviceInfo)
	return nil
}

func (me *meHomeLight) Close() error {
	fmt.Printf("miHomeLight.Close, %v\n", &me.DeviceInfo)
	return nil
}

func (me *meHomeLight) SetLightMode(mode int) error {
	fmt.Printf("miHomeLight.SetLightMode, %v\n", mode)
	me.xMode = mode
	return nil
}

func (me *meHomeLight) GetLightMode() (error, int) {
	fmt.Printf("miHomeLight.GetLightMode, %v\n", me.xMode)
	return nil, me.xMode
}

type meHomeAirConditioner struct {
	DeviceInfo
	temperature float64
}

func newMeHomeAirConditioner(info *DeviceInfo) *meHomeAirConditioner {
	return &meHomeAirConditioner{
		DeviceInfo:  *info,
		temperature: 0,
	}
}

func (me *meHomeAirConditioner) Open() error {
	fmt.Printf("meHomeAirConditioner.open, %v\n", &me.DeviceInfo)
	return nil
}

func (me *meHomeAirConditioner) Close() error {
	fmt.Printf("meHomeAirConditioner.Close, %v\n", &me.DeviceInfo)
	return nil
}

func (me *meHomeAirConditioner) SetTemperature(t float64) error {
	fmt.Printf("meHomeAirConditioner.SetTemperature, %v\n", t)
	me.temperature = t
	return nil
}

func (me *meHomeAirConditioner) GetTemperature() (error, float64) {
	fmt.Printf("meHomeAirConditioner.GetTemperature, %v\n", me.temperature)
	return nil, me.temperature
}

