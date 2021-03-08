package simple_factory

import "testing"

func Test_Simplefactory(t *testing.T) {
	config := make([]*LightInfo, 0)
	config = append(config, newLightInfo(1, "客厅灯", "MIHOME", "L-100"))
	config = append(config, newLightInfo(2, "卧室灯", "REDME", "L-56"))

	factory := defaultLightFactory
	for _, info := range config {
		err, light := factory.Create(info)
		if err != nil {
			t.Errorf(err.Error())
		} else {
			_ = light.Open()
			_ = light.Close()
		}
	}
}
