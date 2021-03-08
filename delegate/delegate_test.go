package delegate

import (
	"fmt"
	"testing"
)

func Test_Delegate(t *testing.T) {
	dispatcher := globalMsgDispatcher
	vEchoMsg := newEchoMsg("msg-1", "echo msg")
	response := dispatcher.Handle(vEchoMsg)
	fmt.Printf("echo response: id=%v, cls=%v, content=%v\n", response.ID(), response.Class(), response.Content())

	vTimeMsg := newTimeMsg("msg-2")
	response = dispatcher.Handle(vTimeMsg)
	fmt.Printf("time response: id=%v, cls=%v, content=%v\n", response.ID(), response.Class(), response.Content())
}
