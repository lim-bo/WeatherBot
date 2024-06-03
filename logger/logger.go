package logger

import (
	"log"
	"os"
)

var (
	file, _ = os.OpenFile("../logs/errors.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	logFile = log.New(file, "", 0)
)

func LogFatalError(err error) {
	if err != nil {
		log.Println(err)
		logFile.Fatalln(err)
	}
}

func LogError(err error) {
	logFile.Println(err)
}

func ClearLogs() {
	logFile.Writer().Write([]byte{})
}
