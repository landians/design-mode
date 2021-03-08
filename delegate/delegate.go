package delegate

import (
	"fmt"
	"reflect"
	"time"
)

// 消息的接口
type IMsg interface {
	ID() string
	Class() string
	Content() string
}

// 消息的基类，实现 IMsg 接口
type baseMsg struct {
	xID      string
	xClass   string
	sContent string
}

func newBaseMsg(id string, cls string, content string) *baseMsg {
	return &baseMsg{
		xID:      id,
		xClass:   cls,
		sContent: content,
	}
}

func (b *baseMsg) ID() string {
	return b.xID
}

func (b *baseMsg) Class() string {
	return b.xClass
}

func (b *baseMsg) Content() string {
	return b.sContent
}

// 继承 baseMsg，实现 PING/PONG 心跳的心跳信息
type echoMsg struct {
	baseMsg
}

func newEchoMsg(id string, content string) *echoMsg {
	return &echoMsg{
		baseMsg: *newBaseMsg(id, "echoMsg", content),
	}
}

// 继承 baseMsg， 用于获取服务器时间的消息
type timeMsg struct {
	baseMsg
}

func newTimeMsg(id string) *timeMsg {
	return &timeMsg{
		baseMsg: *newBaseMsg(id, "timeMsg", time.Now().Format("2006-01-02 15:04:05")),
	}
}


// 消息处理的接口
type IMsgHandler interface {
	Handle(request IMsg) IMsg
}

// 全局消息调度器，用于注册消息处理器，按照类型来分发消息，实现 IMsHandler 接口
type msgDispatchDelegate struct {
	subHandlers map[string]IMsgHandler
}

func (m *msgDispatchDelegate) Register(cls string, handler IMsgHandler) {
	m.subHandlers[cls] = handler
}

func newMsgDispatchDelegate() IMsgHandler {
	it := &msgDispatchDelegate{
		subHandlers: make(map[string]IMsgHandler),
	}

	it.Register("echoMsg", newEchoMsgHandler())
	it.Register("timeMsg", newTimeMsgHandler())

	return it
}

func (m *msgDispatchDelegate) Handle(request IMsg) IMsg {
	if request == nil {
		return nil
	}

	handler, ok := m.subHandlers[request.Class()]
	if !ok {
		fmt.Printf("msgDispatchDelegate.Handle, handler not found: id=%v, cls=%v\n", request.ID(), request.Class())
		return nil
	}

	fmt.Printf("tMsgDispatchDelegate.Handle, handler=%v, id=%v, cls=%v\n", reflect.TypeOf(handler).String(), request.ID(), request.Class())
	return handler.Handle(request)
}

var globalMsgDispatcher = newMsgDispatchDelegate()

// echoMsg 的消息处理器, 实现 IMsgHandler 接口
type echoMsgHandler struct {
}

func newEchoMsgHandler() IMsgHandler {
	return &echoMsgHandler{}
}

func (e *echoMsgHandler) Handle(request IMsg) IMsg {
	fmt.Printf("echoMsgHandler.Handle, id=%v, cls=%v\n", request.ID(), request.Class())
	return request
}

// timeMsg 的消息处理器，实现 IMsgHandler 接口
type timeMsgHandler struct {
}

func newTimeMsgHandler() IMsgHandler {
	return &timeMsgHandler{}
}

func (e *timeMsgHandler) Handle(request IMsg) IMsg {
	fmt.Printf("timeMsgHandler.Handle, id=%v, cls=%v\n", request.ID(), request.Class())
	return request
}
