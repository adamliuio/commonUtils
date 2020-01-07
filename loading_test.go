package commonUtils

import (
	"log"
	"testing"
)

func _TestLoadEnv(t *testing.T) {
	env := LoadEnv(".env")
	_ = env
}

func _TestLoadEnvVar(t *testing.T) {
	ok := LoadEnvVar(".env", "TWO")
	log.Println("ok:", ok)
}
