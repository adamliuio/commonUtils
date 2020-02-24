package commonUtils

import (
	"log"
	"testing"
)

func TestLoadEnv(t *testing.T) {
	env, err := LoadEnv(".env")
	if err != nil {
		return
	}
	_ = env
}

func TestLoadEnvVar(t *testing.T) {
	ok, err := LoadEnvVar(".env", "TWO")
	if err != nil {
		return
	}
	log.Println("ok:", ok)
}
