package proxy

import (
	"errors"
	"fmt"
)

// 系统用户接口，用于订单服务上下文
type IUser interface {
	ID() int
	Name() string
	Allowed(perm string) bool
}

// 模拟用户数据，封装运行时的用户信息
type mockUser struct {
	xID          int
	xName        string
	xPermissions map[string]bool
}

func newMockUser(id int, name string, perms []string) *mockUser {
	it := &mockUser{
		xID:          id,
		xName:        name,
		xPermissions: make(map[string]bool, len(perms)),
	}

	for _, k := range perms {
		it.xPermissions[k] = true
	}
	return it
}

func (m *mockUser) ID() int {
	return m.xID
}

func (m *mockUser) Name() string {
	return m.xName
}

func (m *mockUser) Allowed(perm string) bool {
	if m.xPermissions == nil {
		return false
	}

	_, ok := m.xPermissions[perm]
	return ok
}

// 订单信息
type Order struct {
	ID             int
	OrderNo        string
	CustomerID     int
	OrderDate      string
	ReceiveAddress string
}

// 订单服务接口
type IOrderService interface {
	Load(id int) (error, *Order)
	Save(order *Order, user IUser) error
	Delete(id int, user IUser) error
}

// 模拟订单服务数据，实现 IOrderService 接口，提供订单的基本增删改查功能
type mockOrderService struct {
	items map[int]*Order
}

func newMockOrderService() IOrderService {
	return &mockOrderService{
		items: make(map[int]*Order),
	}
}

func (m *mockOrderService) Load(id int) (error, *Order) {
	it, ok := m.items[id]
	if ok {
		return nil, it
	} else {
		return errors.New("no such order"), nil
	}
}

func (m *mockOrderService) Save(o *Order, user IUser) error {
	m.items[o.ID] = o
	return nil
}

func (m *mockOrderService) Delete(id int, user IUser) error {
	_, ok := m.items[id]
	if ok {
		delete(m.items, id)
	} else {
		return errors.New("no such order")
	}
	return nil
}

// 订单服务代理，以代理模式来增加订单的 Save 和 Delete 日志
type proxyOrderServiceLog struct {
	orderService IOrderService
}

func newProxyOrderServiceLog(service IOrderService) IOrderService {
	return &proxyOrderServiceLog{orderService: service}
}

func (p *proxyOrderServiceLog) Load(id int) (error, *Order) {
	return p.orderService.Load(id)
}

func (p *proxyOrderServiceLog) Save(o *Order, user IUser) error {
	err := p.orderService.Save(o, user)
	fmt.Printf("IOrderService.Save, user=%v, order=%v, error=%v\n", user.Name(), o, err)
	return err
}

func (p *proxyOrderServiceLog) Delete(id int, user IUser) error {
	err := p.orderService.Delete(id, user)
	fmt.Printf("IOrderService.Delete, user=%v, order.id=%v, error=%v\n", user.Name(), id, err)
	return err
}

// 订单服务代理，以代理模式来校验订单的安全
type proxyOrderServiceSecure struct {
	orderService IOrderService
}

func newProxyOrderServiceSecure(service IOrderService) *proxyOrderServiceSecure {
	return &proxyOrderServiceSecure{orderService: service}
}

func (p *proxyOrderServiceSecure) Load(id int) (error, *Order) {
	return p.orderService.Load(id)
}

func (p *proxyOrderServiceSecure) Save(o *Order, user IUser) error {
	if !user.Allowed("order.save") {
		return errors.New("permission denied")
	}
	return p.orderService.Save(o, user)
}

func (p *proxyOrderServiceSecure) Delete(id int, user IUser) error {
	if !user.Allowed("order.delete") {
		return errors.New("permission denied")
	}
	return p.orderService.Delete(id, user)
}
