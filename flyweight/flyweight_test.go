package flyweight

import "testing"

func Test_Flyweight(t *testing.T) {
	from := "福田"
	to := "广州南"
	ticket := newMockTicket(1, from, to, 100)
	mockTicketService.Save(ticket)
	mockTicketRemainingService.Save(ticket.ID(), 10)

	remaining := mockTicketRemainingService.Get(from, to)
	t.Logf("from=%s, to=%s, price=%v, remaining=%v\n", remaining.From(), remaining.To(), remaining.Price(), remaining.Remaining())
}
