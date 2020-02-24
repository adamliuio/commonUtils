package commonUtils

import (
	"testing"
)

var url string = ""

func TestServiceStatusUpdate(t *testing.T) {
	a := map[string]string{}
	a["title"] = "this is a title"
	a["value"] = "and here's the value"
	arr := []map[string]string{}
	arr = append(arr, a)

	ServiceStatusUpdate("Useful-Tools", arr, url)
}

func TestNotifyMaster(t *testing.T) {
	NotifyMaster("god damn!", url)
}
