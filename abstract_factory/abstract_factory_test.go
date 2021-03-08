package abstract_factory

import "testing"

func Test_AbstractFactory(t *testing.T) {
	// 注册不同厂商实现的智能设备工厂
	defaultFactoryRegistry.Set("MIHOME", newMeHomeFactory())

	config := make([]*DeviceInfo, 0)
	config = append(config, NewDeviceInfo(1, "客厅灯", DeviceTypeLight, "MIHOME", "L-100"))
	config = append(config, NewDeviceInfo(2, "主卧空调", DeviceTypeAirConditioner, "MIHOME", "GeLi"))

	for _, info := range config {
		factory := defaultFactoryRegistry.Get(info.Vendor())
		if factory == nil {
			t.Errorf("unsupported vendor: %s", info.Vendor())
		} else {
			switch info.DeviceType() {
			case DeviceTypeLight:
				err, light := factory.CreateLight(info)
				if err != nil {
					t.Error(err.Error())
				} else {
					_ = light.Open()
					_ = light.SetLightMode(1)
					_, _ = light.GetLightMode()
					_ = light.Close()
				}
			case DeviceTypeAirConditioner:
				err, air := factory.CreateAirConditioner(info)
				if err != nil {
					t.Error(err.Error())
				} else {
					_ = air.Open()
					_ = air.SetTemperature(12.0)
					_, _ = air.GetTemperature()
					_ = air.Close()
				}
			default:
				t.Errorf("unsupported deviceType :%v\n", info.DeviceType())
			}
		}
	}
}
