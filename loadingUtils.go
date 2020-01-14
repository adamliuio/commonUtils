package commonUtils

import (
	"github.com/joho/godotenv"

	"fmt"
	"log"
	"path/filepath"
	"runtime"
)

func LoadEnv(envFilePath string) (env map[string]string) {

	// ::envFilePath:: root should be where you execute the program from
	var err error
	env, err = godotenv.Read(envFilePath)
	if err != nil {
		pc, file, line, _ := runtime.Caller(1) // using 1 to log where the error happened
		file = filepath.Base(file)
		n := runtime.FuncForPC(pc).Name()
		var msg = fmt.Sprintf("[ERROR] %s (%s:%d)\n%+v\n", n, file, line, err)
		log.Fatalln(msg)
	}
	return
}

func LoadEnvVar(envFilePath, varName string) (v string) {

	// ::envFilePath:: root should be where you execute the program from
	env, err := godotenv.Read(envFilePath)
	if err != nil {
		pc, file, line, _ := runtime.Caller(1) // using 1 to log where the error happened
		file = filepath.Base(file)
		n := runtime.FuncForPC(pc).Name()
		var msg = fmt.Sprintf("[ERROR] %s (%s:%d)\n%+v\n", n, file, line, err)
		log.Fatalln(msg)
	}
	v = env[varName]
	return
}
