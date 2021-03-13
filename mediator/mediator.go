package mediator

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"sync"
)

// 模拟手机APP，用于和云服务中心通信，控制智能设备
type mockPhoneApp struct {
	mediator ICloudMediator
}

func newMockPhoneApp(mediator ICloudMediator) *mockPhoneApp {
	return &mockPhoneApp{mediator: mediator}
}

func (m *mockPhoneApp) LightOpen(id int) error {
	return m.lightCommand(id, "light.open")
}

func (m *mockPhoneApp) LightClose(id int) error {
	return m.lightCommand(id, "light.close")
}

func (m *mockPhoneApp) lightCommand(id int, cmd string) error {
	res := m.mediator.Command(id, cmd)
	if res != "OK" {
		return errors.New(res)
	}
	return nil
}

// 云服务中心面向 App 端的接口
type ICloudMediator interface {
	Command(id int, cmd string) string
}

// 云服务中心面向智能设备的注册接口
type ICloudCenter interface {
	Register(dev ISmartDevice)
}

// 智能设备接口
type ISmartDevice interface {
	ID() int
	Command(cmd string) string
}

// 模拟云服务中心
type xMockCloudMediator struct {
	devices map[int] ISmartDevice
	mu sync.RWMutex
}

func newMockCloudMediator() ICloudMediator {
	return &xMockCloudMediator{
		devices: make(map[int]ISmartDevice),
	}
}

func (m *xMockCloudMediator) Command(id int, cmd string) string {
	m.mu.RLock()
	defer m.mu.RUnlock()

	it,ok := m.devices[id]
	if !ok {
		return "device not found"
	}
	return it.Command(cmd)
}

func (m *xMockCloudMediator) Register(dev ISmartDevice) {
	m.mu.Lock()
	m.devices[dev.ID()] = dev
	m.mu.Unlock()
}

var defaultCloudMediator = newMockCloudMediator()
var defaultCloudCenter = defaultCloudMediator.(ICloudCenter)

// 模拟智能灯设备
type xMockSmartLight struct {
	id int
}

func newMockSmartLight(id int) ISmartDevice {
	return &xMockSmartLight{id: id}
}

func (m *xMockSmartLight) ID() int {
	return m.id
}

func (m *xMockSmartLight) Command(cmd string) string {
	if cmd == "light open" {
		err := m.open()
		if err != nil {
			return err.Error()
		}
	} else if cmd == "light close" {
		err := m.close()
		if err != nil {
			return err.Error()
		}
	} else if strings.HasPrefix(cmd, "light switch_mode") {
		args := strings.Split(cmd, " ")
		if len(args) != 3 {
			return "invalid switch command"
		}

		n, err := strconv.Atoi(args[2])
		if err != nil {
			return "invalid mode number"
		}

		err = m.switchMode(n)
		if err != nil {
			return err.Error()
		}

	} else {
		return "unrecognized command"
	}

	return "OK"
}

func (m *xMockSmartLight) open() error {
	fmt.Printf("xMockSmartLight.open, id=%v\n", m.id)
	return nil
}

func (m *xMockSmartLight) close() error {
	fmt.Printf("xMockSmartLight.close, id=%v\n", m.id)
	return nil
}

func (m *xMockSmartLight) switchMode(mode int) error {
	fmt.Printf("xMockSmartLight.switchMode, id=%v, mode=%v\n", m.id, mode)
	return nil
}
