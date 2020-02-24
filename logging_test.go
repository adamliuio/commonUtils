package commonUtils

import (
	"errors"
	"testing"
)

func TestLogError(t *testing.T) {
	var err = errors.New("damn")
	LogError(err)
}

func TestLogString(t *testing.T) {
	LogString("damn err")
}
