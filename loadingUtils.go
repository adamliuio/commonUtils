package commonUtils

import (
	"github.com/joho/godotenv"

	"log"
)

func LoadEnv(envFilePath string) (env map[string]string) {
	var err error
	env, err = godotenv.Read(envFilePath)
	if err != nil {
		log.Fatalln(err)
	}
	return
}

func LoadEnvVar(envFilePath, varName string) (v string) {
	env := LoadEnv(envFilePath)
	v = env[varName]
	return
}
