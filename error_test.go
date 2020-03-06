package commonUtils

import (
	"errors"
	"testing"
)

func _TestCheckError(t *testing.T) {
	err := errors.New("damn")
	CheckError(err, "debug", "")
}
