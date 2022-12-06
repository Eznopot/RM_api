package logger

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"sync"
	"time"

	"gopkg.in/yaml.v2"
)

type LogConfig struct {
	Logger struct {
		Path string `yaml:"path"`
	} `yaml:"logger"`
}

var once sync.Once

var (
	config LogConfig
)

func Init() {
	once.Do(
		func() {
			f, err := os.Open("config.yml")
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(-1)
			}
			defer f.Close()
			decoder := yaml.NewDecoder(f)
			err = decoder.Decode(&config)
			if err != nil {
				os.Exit(-1)
			}
		},
	)
}

func SetLogFile(path string) *os.File {
	year, month, day := time.Now().Date()
	fileName := fmt.Sprintf("%v-%v-%v.log", day, month.String(), year)
	filePath, err := os.OpenFile(path+"/"+fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	return filePath
}

func Info(message ...string) {
	log.New(SetLogFile(config.Logger.Path), "[Info]: ", log.Lmsgprefix|log.Ldate|log.Ltime|log.Lshortfile).Println(message)
}

func Error(message ...string) {
	log.New(SetLogFile(config.Logger.Path), "[Error]: ", log.Lmsgprefix|log.Ldate|log.Ltime|log.Lshortfile).Println(message, "in", Trace())
}

func Debug(message ...string) {
	log.New(SetLogFile(config.Logger.Path), "[Debug]: ", log.Lmsgprefix|log.Ldate|log.Ltime|log.Lshortfile).Println(message)
}

func Trace() (string) {
	pc, _, _, ok := runtime.Caller(2)
	if !ok {
		return "?"
	}

	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return "?"
	}

	return fn.Name()
}
