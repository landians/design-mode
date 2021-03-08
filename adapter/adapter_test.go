package adapter

import "testing"

func Test_Adapter(t *testing.T) {
	factory := specialThermometerFactory
	thermometer := factory.Create("some configuration")
	t.Logf("centigrade = %v", thermometer.Centigrade())
}