package commonUtils

import (
	"errors"
	"time"
)

type Function func()

func RunEveryDayAt(f Function, hour, minute, second, milisecond int) {
	t := time.Now()
	shoudRunAt := time.Date(t.Year(), t.Month(), t.Day(), hour, minute, second, milisecond, t.Location())
	duration := time.Hour * 24
	run(f, shoudRunAt, duration)
}

func RunEveryHourAt(f Function, minute, second, milisecond int) {
	t := time.Now()
	shoudRunAt := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), minute, second, milisecond, t.Location())
	duration := time.Hour
	run(f, shoudRunAt, duration)
}

func RunEveryMinuteAt(f Function, second, milisecond int) {
	t := time.Now()
	shoudRunAt := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), second, milisecond, t.Location())
	duration := time.Minute
	run(f, shoudRunAt, duration)
}

func run(f Function, shoudRunAt time.Time, duration time.Duration) {
	gap := shoudRunAt.Sub(time.Now())
	if gap < 0 {
		shoudRunAt = shoudRunAt.Add(duration)
		gap = shoudRunAt.Sub(time.Now()) // using time.Now() here to reduce the time difference to the minimum
	}
	for {
		time.Sleep(gap)
		gap = duration
		f()
	}
}

func RunEveryNMinute(f Function, n int64, runAtStart bool) {
	duration := time.Minute * time.Duration(n)
	if runAtStart {
		f()
	}
	for {
		time.Sleep(duration)
		f()
	}
}

func RunEveryNSecond(f Function, n int64, runAtStart bool) {
	duration := time.Second * time.Duration(n)
	if runAtStart {
		f()
	}
	for {
		time.Sleep(duration)
		f()
	}
}

func RunEvery(f Function, n int, timeUnit string, runAtStart bool) error {
	var duration time.Duration
	switch timeUnit {
	case "Day":
		duration = time.Hour * time.Duration(n*24)
	case "Hour":
		duration = time.Hour * time.Duration(n)
	case "Minute":
		duration = time.Minute * time.Duration(n)
	case "Second":
		duration = time.Second * time.Duration(n)
	default:
		return errors.New("timeUnit can only be one of 'Day', 'Hour', 'Minute', 'Second'.")
	}
	if runAtStart {
		f()
	}
	for {
		time.Sleep(duration)
		f()
	}
	return nil
}
