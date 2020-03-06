package commonUtils

import (
	"log"
	"testing"
)

func _TestLoadEnv(t *testing.T) {
	env, err := LoadEnv(".env")
	if err != nil {
		return
	}
	_ = env
}

func _TestLoadEnvVar(t *testing.T) {
	ok, err := LoadEnvVar(".env", "TWO")
	if err != nil {
		return
	}
	log.Println("ok:", ok)
}
