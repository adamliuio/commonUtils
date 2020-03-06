package commonUtils

import (
	"testing"
)

func _TestGetFileAndLine(t *testing.T) {
	dir, file, line := GetDirFileLine()
	t.Log("dir:", dir)
	t.Log("file:", file)
	t.Log("line:", line)
}
