package commonUtils

import (
	"time"
)

type Function func()

func RunEveryDayAt(f Function, hour, minute, second, milisecond int) {
	t := time.Now()
	n := time.Date(t.Year(), t.Month(), t.Day(), hour, minute, second, milisecond, t.Location())
	d := n.Sub(t)
	if d < 0 {
		n = n.Add(24 * time.Hour)
		d = n.Sub(t)
	}
	for {
		time.Sleep(d)
		d = 24 * time.Hour
		f()
		// n = time.Date(t.Year(), t.Month(), t.Day(), hour, minute, second, milisecond, t.Location())
	}
}

func RunEveryHourAt(f Function, minute, second, milisecond int) {
	t := time.Now()
	n := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), minute, second, milisecond, t.Location())
	d := n.Sub(t)
	if d < 0 {
		n = n.Add(time.Hour)
		d = n.Sub(t)
	}
	for {
		time.Sleep(d)
		d = time.Hour
		f()
		// n = time.Date(t.Year(), t.Month(), t.Day(), hour, minute, second, milisecond, t.Location())
	}
}
