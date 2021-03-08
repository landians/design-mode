package flyweight

import (
	"strings"
	"sync"
)

// 车票的基本信息接口
type ITicket interface {
	ID() int
	From() string
	To() string
	LeavingTime() string
	ArrivalTime() string
	InterList() []string
	Price() float64
}

// 余票信息接口，继承了车票信息的基本接口
type ITicketRemaining interface {
	ITicket
	Remaining() int
}

// 车票信息服务接口
type ITicketService interface {
	Get(from string, to string) ITicket
	Save(it ITicket)
}

// 余票信息服务接口，根据发站和到站，查询余票信息
type ITicketRemainingService interface {
	Get(from string, to string) ITicketRemaining
	Save(id int, num int)
}

// mock 车票信息实体
type mockTicket struct {
	xID          int
	xFrom        string
	xTo          string
	xLeavingTime string
	xArrivalTime string
	xInterList   []string
	xPrice       float64
}

func newMockTicket(id int, from string, to string, price float64) *mockTicket {
	return &mockTicket{
		xID:          id,
		xFrom:        from,
		xTo:          to,
		xLeavingTime: "09:00",
		xArrivalTime: "11:30",
		xInterList:   strings.Split("深圳北,虎门", ","),
		xPrice:       price,
	}
}

func (m *mockTicket) ID() int {
	return m.xID
}

func (m *mockTicket) From() string {
	return m.xFrom
}

func (m *mockTicket) To() string {
	return m.xTo
}

func (m *mockTicket) LeavingTime() string {
	return m.xLeavingTime
}

func (m *mockTicket) ArrivalTime() string {
	return m.xArrivalTime
}

func (m *mockTicket) InterList() []string {
	return m.xInterList
}

func (m *mockTicket) Price() float64 {
	return m.xPrice
}

// mock 车票信息服务实体，通过享元模式池化了车票信息
type xMockTicketService struct {
	tickets map[string]ITicket			// 车票的池子
	mu sync.RWMutex
}

func newMockTicketService() *xMockTicketService {
	return &xMockTicketService{
		tickets: make(map[string]ITicket, 16),
	}
}

func (ms *xMockTicketService) Get(from string, to string) ITicket {
	k := from + " - " + to

	ms.mu.RLock()
	defer ms.mu.RUnlock()
	it, ok := ms.tickets[k]
	if ok {
		return it
	}
	return nil
}

func (ms *xMockTicketService) Save(it ITicket) {
	k := it.From() + " - " + it.To()

	ms.mu.Lock()
	ms.tickets[k] = it
	ms.mu.Unlock()
}

var mockTicketService ITicketService = newMockTicketService()

// 余票信息实体
type xMockTicketRemaining struct {
	ITicket
	remaining int
}

func newMockTicketRemaining(it ITicket, num int) *xMockTicketRemaining {
	return &xMockTicketRemaining{
		ITicket:   it,
		remaining: num,
	}
}

func (mr *xMockTicketRemaining) Remaining() int {
	return mr.remaining
}

// 余票信息服务实体, 根据发站和到站，查询余票信息
type xMockTicketRemainingService struct {
	remaining map[int]int
	mu sync.RWMutex
}

func newMockTicketRemainingService() *xMockTicketRemainingService {
	return &xMockTicketRemainingService{
		remaining: make(map[int]int, 16),
	}
}

func (mrs *xMockTicketRemainingService) Get(from string, to string) ITicketRemaining {
	ticket := mockTicketService.Get(from, to)
	if ticket == nil {
		return nil
	}

	r := newMockTicketRemaining(ticket, 0)
	mrs.mu.RLock()
	defer mrs.mu.RUnlock()
	num, ok := mrs.remaining[ticket.ID()]
	if ok {
		r.remaining = num
	}
	return r
}

func (mrs *xMockTicketRemainingService) Save(id int, num int) {
	mrs.mu.Lock()
	defer mrs.mu.Unlock()
	mrs.remaining[id] = num
}

var mockTicketRemainingService ITicketRemainingService = newMockTicketRemainingService()