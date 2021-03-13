package visitor

import "testing"

func Test_Visitor(t *testing.T) {
	service := mockSaleOrderService

	_ = service.Save(newSaleOrder(1, "张三", "广州", "电视", 10))
	_ = service.Save(newSaleOrder(2, "李四", "深圳", "冰箱", 20))
	_ = service.Save(newSaleOrder(3, "王五", "东莞", "空调", 30))
	_ = service.Save(newSaleOrder(4, "张三三", "广州", "空调", 10))
	_ = service.Save(newSaleOrder(5, "李四四", "深圳", "电视", 20))
	_ = service.Save(newSaleOrder(6, "王五五", "东莞", "冰箱", 30))

	// test CityVisitor
	cv := newCityVisitor()
	service.Visit(cv)
	cv.Report()

	// test ProductVisitor
	pv := newProductVisitor()
	service.Visit(pv)
	pv.Report()
}