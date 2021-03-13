package visitor

import "fmt"

// 销售订单实体
type SaleOrder struct {
	ID       int
	Customer string
	City     string
	Product  string
	Quantity int
}

func newSaleOrder(id int, customer string, city string, product string, quantity int) *SaleOrder {
	return &SaleOrder{
		ID:       id,
		Customer: customer,
		City:     city,
		Product:  product,
		Quantity: quantity,
	}
}

// 销售订单服务接口
type ISaleOrderService interface {
	Save(order *SaleOrder) error
	Visit(visitor ISaleOrderVisitor)
}

// 销售订单访问者接口
type ISaleOrderVisitor interface {
	Visit(it *SaleOrder)
	Report()
}

// 模拟销售订单服务
type xMockSaleOrderService struct {
	orders map[int]*SaleOrder
}

func newMockSaleOrderService() ISaleOrderService {
	return &xMockSaleOrderService{
		orders: make(map[int]*SaleOrder, 0),
	}
}

func (m *xMockSaleOrderService) Save(it *SaleOrder) error {
	m.orders[it.ID] = it
	return nil
}

func (m *xMockSaleOrderService) Visit(visitor ISaleOrderVisitor) {
	for _, v := range m.orders {
		visitor.Visit(v)
	}
}

var mockSaleOrderService = newMockSaleOrderService()

// 区域销售报表, 按城市汇总销售情况, 实现ISaleOrderVisitor接口
type cityVisitor struct {
	cities map[string]int
}

func newCityVisitor() ISaleOrderVisitor {
	return &cityVisitor{
		cities: make(map[string]int, 0),
	}
}

func (c *cityVisitor) Visit(it *SaleOrder) {
	n, ok := c.cities[it.City]
	if ok {
		c.cities[it.City] = n + it.Quantity
	} else {
		c.cities[it.City] = it.Quantity
	}
}

func (c *cityVisitor) Report() {
	for k, v := range c.cities {
		fmt.Printf("city=%s, sum=%v\n", k, v)
	}
}

// 品类销售报表, 按产品汇总销售情况, 实现ISaleOrderVisitor接口
type productVisitor struct {
	products map[string]int
}

func newProductVisitor() ISaleOrderVisitor {
	return &productVisitor{
		products: make(map[string]int, 0),
	}
}

func (p *productVisitor) Visit(it *SaleOrder) {
	n, ok := p.products[it.Product]
	if ok {
		p.products[it.Product] = n + it.Quantity
	} else {
		p.products[it.Product] = it.Quantity
	}
}

func (p *productVisitor) Report() {
	for k, v := range p.products {
		fmt.Printf("product=%s, sum=%v\n", k, v)
	}
}
