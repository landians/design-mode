package adapter

// 现有系统的温度计接口，单位: 摄氏度
type IThermometer interface {
	Centigrade() float64
}

// 现有系统的温度计工厂接口
type IThermometerFactory interface {
	Create(config string) IThermometer
}

// 厂家 sdk 提供的温度计接口, 单位: 华氏度
type ISpecialThermometer interface {
	Fahrenheit() float64
}

// 厂家 sdk 提供的温度计实现
type mockSpecialThermometer struct {
	address string
}

func newMockSpecialThermometer(address string) ISpecialThermometer {
	return &mockSpecialThermometer{address: address}
}

func (m *mockSpecialThermometer) Fahrenheit() float64 {
	return 79.7
}

// 适配厂家 sdk 提供的温度计
type mockSpecialAdapter struct {
	origin ISpecialThermometer
}

func newMockSpecialAdapter(origin ISpecialThermometer) IThermometer {
	return &mockSpecialAdapter{origin: origin}
}

// 适配实现
func (ms *mockSpecialAdapter) Centigrade() float64 {
	return (ms.origin.Fahrenheit() - 32) * 5 / 9
}

// 新温度计的工厂类，实现 IThermometerFactory 接口
type mockSpecialFactory struct {
}

func newMockSpecialFactory() IThermometerFactory {
	return &mockSpecialFactory{}
}

func (msf *mockSpecialFactory) Create(config string) IThermometer {
	t := newMockSpecialThermometer(msf.parseAddress(config))
	return newMockSpecialAdapter(t)
}

func (msf *mockSpecialFactory) parseAddress(config string) string{
	return "http://localhost:8080"
}

var specialThermometerFactory IThermometerFactory = newMockSpecialFactory()