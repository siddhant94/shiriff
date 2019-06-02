package logging

import (
	"os"
	"log"
	"shiriff/config"
)
var logFilePath string
func init () {
	logFilePath = config.LOGPATH
}
func Error (msg string, err error) {
	f, err := os.OpenFile(logFilePath,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	logger := log.New(f, "Error", log.LstdFlags)
	logger.Println(msg)
	logger.Println(err)
}