package observer

import (
	"testing"
	"time"
)

func Test_Observer(t *testing.T) {
	_ = newAlarmClock("下午开会", 14, 30, false)

	_ = newAlarmClock("起床", 6, 0, true)
	_ = newAlarmClock("午饭", 12, 30, true)
	_ = newAlarmClock("午休", 13, 0, true)
	_ = newAlarmClock("晚饭", 18, 30,  true)
	clock := newAlarmClock("晚安", 22, 0, true)

	for {
		if clock.Occurs() >= 2 {
			break
		}
		time.Sleep(1 * time.Second)
	}
}
