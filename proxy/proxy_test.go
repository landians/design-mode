package proxy

import (
	"strings"
	"testing"
	"time"
)

func Test_Proxy(t *testing.T) {
	admin := newMockUser(1, "管理员", strings.Split("order.load,order.save,order.delete", ","))
	guest := newMockUser(2, "游客", strings.Split("order.load", ","))

	order := &Order{
		ID:             1,
		OrderNo:        "mock-order-1",
		CustomerID:     1,
		OrderDate:      time.Now().Format("2006-01-02"),
		ReceiveAddress: "mock-address",
	}

	os1 := newMockOrderService()
	fnCallAndLog := func(fn func() error) {
		err := fn()
		if err != nil {
			t.Log(err)
		}
	}
	fnCallAndLog(func() error {
		return os1.Save(order, admin)
	})
	fnCallAndLog(func() error {
		return os1.Delete(order.ID, admin)
	})

	// os2 代理 os1
	os2 := newProxyOrderServiceLog(os1)
	fnCallAndLog(func() error {
		return os2.Save(order, admin)
	})
	fnCallAndLog(func() error {
		return os2.Delete(order.ID, admin)
	})

	// os3 -> 代理 -> (osx -> os1)
	os3 := newProxyOrderServiceLog(newProxyOrderServiceSecure(os1))
	fnCallAndLog(func() error {
		return os3.Save(order, admin)
	})
	fnCallAndLog(func() error {
		return os3.Delete(order.ID, admin)
	})
	fnCallAndLog(func() error {
		return os3.Save(order, guest)
	})
	fnCallAndLog(func() error {
		return os3.Delete(order.ID, guest)
	})
}
