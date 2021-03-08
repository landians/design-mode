package factory_method

import "testing"

func Test_FactoryMethod(t *testing.T) {
	// 注册不同厂商的智能灯生产工厂
	defaultFactoryRegistry.Set("MIHOME", newMeHomeLightFactory())
	defaultFactoryRegistry.Set("REDME", newRedMeLightFactory())

	config := make([]*LightInfo, 0)
	config = append(config, newLightInfo(1, "客厅灯", "MIHOME", "L-100"))
	config = append(config, newLightInfo(2, "卧室灯", "REDME", "L-56"))

	for _, info := range config {
		// 获取这个厂商的智能灯工厂
		factory := defaultFactoryRegistry.Get(info.Vendor())
		if factory == nil {
			t.Errorf("unsupported vendor: %s", info.Vendor())
		} else {
			// 再由对应的智能灯工厂开始生产智能灯
			err, light := factory.Create(info)
			if err != nil {
				t.Error(err.Error())
			} else {
				_ = light.Open()
				_ = light.Close()
			}
		}
	}
}
