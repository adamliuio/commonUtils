package commonUtils

import (
	"testing"
)

func TestRoutine(t *testing.T) {
	dir, file, line := GetFileAndLine()
	t.Log("dir:", dir)
	t.Log("file:", file)
	t.Log("line:", line)
}
