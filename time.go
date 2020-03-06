package commonUtils

import (
	"time"
)

func TimeStringToUnix(ts string) (unix int64, err error) {
	var t time.Time
	t, err = time.Parse(time.RFC3339, ts)
	if err != nil {
		return
	}
	unix = t.Unix()
	return
}

func UnixTimestampToString(unix int64) (ts string) {
	var tm time.Time = time.Unix(unix, 0)
	ts = tm.Format(time.RFC3339)
	return
}

func TimeStringToTime(ts string) (t time.Time, err error) {
	t, err = time.Parse(time.RFC3339, ts)
	if err != nil {
		return
	}
	return
}

func TimeNowUnix() int64 {
	return time.Now().Unix()
}

func TimeNowStr() string {
	return time.Now().Format(time.RFC3339)
}
