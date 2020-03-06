package commonUtils

import (
	"github.com/joho/godotenv"

	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func LoadEnv(envFilePath string) (env map[string]string, err error) {
	_, err = os.Stat(envFilePath)
	if err != nil {
		return
	}

	// ::envFilePath:: root should be where you execute the program from
	env, err = godotenv.Read(envFilePath)
	if err != nil {
		pc, file, line, _ := runtime.Caller(1) // using 1 to log where the error happened
		file = filepath.Base(file)
		n := runtime.FuncForPC(pc).Name()
		var msg = fmt.Sprintf("[ERROR] %s (%s:%d)\n%+v\n", n, file, line, err)
		err = errors.New(msg)
	}
	return
}

func LoadEnvVar(envFilePath, varName string) (v string, err error) {

	// ::envFilePath:: root should be where you execute the program from
	var env map[string]string
	env, err = godotenv.Read(envFilePath)
	if err != nil {
		pc, file, line, _ := runtime.Caller(1) // using 1 to log where the error happened
		file = filepath.Base(file)
		n := runtime.FuncForPC(pc).Name()
		var msg = fmt.Sprintf("[ERROR] %s (%s:%d)\n%+v\n", n, file, line, err)
		err = errors.New(msg)
	}
	v = env[varName]
	return
}
