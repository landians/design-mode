package singleton

import "testing"

func Test_Singleton(t *testing.T) {
	// 注册单例
	defaultBeanContainer.SetBean("IDemoSingleton", newContainedSingleton())

	fnTestSingleton := func(it IDemoSingleton)  {
		it.Hello()
	}

	fnTestSingleton(gHungrySingleton)
	fnTestSingleton(GetDoubleCheckSingleton())

	if ok, it := defaultBeanContainer.GetBean("IDemoSingleton"); ok {
		fnTestSingleton(it.(IDemoSingleton))
	}
}
