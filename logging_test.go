package commonUtils

import (
	"errors"
	"testing"
)

func _TestLogError(t *testing.T) {
	var err = errors.New("damn")
	LogError(err)
}

func _TestLogString(t *testing.T) {
	LogString("damn err")
}
