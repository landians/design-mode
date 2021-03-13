package observer

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// 时间服务的接口，接受观察者的注册和注销
type ITimeService interface {
	Attach(observer ITimeObserver)
	Detach(id string)
}

// 时间观察者接口，接收时间变化事件的通知
type ITimeObserver interface {
	ID() string
	TimeElapsed(now *time.Time)
}

// 模拟时间服务，自定义时间倍率以方便时钟相关的测试
type xMockTimeService struct {
	observers map[string]ITimeObserver
	mu        sync.RWMutex
	speed     int64
	state     int64
}

func newMockTimeService(speed int64) ITimeService {
	it := &xMockTimeService{
		observers: make(map[string]ITimeObserver),
		speed:     speed,
	}
	it.Start()
	return it
}

func (m *xMockTimeService) Start() {
	if !atomic.CompareAndSwapInt64(&m.state, 0, 1) {
		return
	}

	go func() {
		timeFrom := time.Now()
		timeOffset := timeFrom.UnixNano()

		for range time.Tick(time.Duration(100) * time.Millisecond) {
			if m.state == 0 {
				break
			}

			nanos := (time.Now().UnixNano() - timeOffset) * m.speed
			t := timeFrom.Add(time.Duration(nanos) * time.Nanosecond)

			m.NotifyAll(&t)
		}
	}()
}

func (m *xMockTimeService) NotifyAll(now *time.Time) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	for _, it := range m.observers {
		go it.TimeElapsed(now)
	}
}

func (m *xMockTimeService) Attach(it ITimeObserver) {
	m.mu.Lock()
	m.observers[it.ID()] = it
	m.mu.Unlock()
}

func (m *xMockTimeService) Detach(id string) {
	m.mu.Lock()
	delete(m.observers, id)
	m.mu.Unlock()

}

var globalTimeService = newMockTimeService(1800)

// 闹铃，实现 ITimeObserver 接口，用于订阅时间变化通知
type alarmClock struct {
	id         string
	name       string
	hour       time.Duration
	minute     time.Duration
	repeatable bool
	next       *time.Time
	occurs     int
}

var gClockID int64 = 0

func newClockID() string {
	id := atomic.AddInt64(&gClockID, 1)
	return fmt.Sprintf("AlarmClock-%d", id)
}

func newAlarmClock(name string, hour int, minute int, repeatable bool) *alarmClock {
	it := &alarmClock{
		id:         newClockID(),
		name:       name,
		hour:       time.Duration(hour),
		minute:     time.Duration(minute),
		repeatable: repeatable,
		next:       nil,
		occurs:     0,
	}

	it.next = it.nextAlarmTime()
	globalTimeService.Attach(it)

	return it
}

func (a *alarmClock) nextAlarmTime() *time.Time {
	now := time.Now()
	today, _ := time.ParseInLocation("2006-01-02 15:04:05", fmt.Sprintf("%s 00:00:00", now.Format("2006-01-02")), time.Local)
	t := today.Add(a.hour * time.Hour).Add(a.minute * time.Minute)
	if t.Unix() < now.Unix() {
		t = t.Add(24 * time.Hour)
	}
	fmt.Printf("%s.next = %s\n", a.name, t.Format("2006-01-02 15:04:05"))
	return &t
}

func (a *alarmClock) ID() string {
	return a.name
}

func (a *alarmClock) TimeElapsed(now *time.Time) {
	it := a.next
	if it == nil {
		return
	}

	if now.Unix() >= it.Unix() {
		a.occurs++
		fmt.Printf("%s 时间=%s 闹铃 %s\n", time.Now().Format("2006-01-02 15:04:05"), now.Format("2006-01-02 15:04:05"), a.name)

		if a.repeatable {
			t := a.next.Add(24 * time.Hour)
			a.next = &t

		} else {
			globalTimeService.Detach(a.ID())
		}
	}
}

func (a *alarmClock) Occurs() int {
	return a.occurs
}
