package logger

import (
	"fmt"
	"os"
	"time"
)

type logger interface {
	log(message string)
}

type (
	consoleLogger struct{}
	fileLogger    struct{ filePath string }
)

func (cl consoleLogger) log(message string) {
	fmt.Println(message)
}

func (fl fileLogger) log(message string) {
	file, err := os.OpenFile(fl.filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening log file:", err)
		return
	}
	defer file.Close()
	logMessage := fmt.Sprintf("%v - %v\n", time.Now().Format(time.RFC3339), message)
	if _, err = file.WriteString(logMessage); err != nil {
		fmt.Println("Error writing to log file:", err)
	}
}

func useLogger(logger logger) {
	logger.log("Running method")
}

func testFunction() {
	cl := consoleLogger{}
	fl := fileLogger{}
	useLogger(cl)
	useLogger(fl)

	var i logger = fileLogger{filePath: "Hello"}
	s := i.(fileLogger)
	fmt.Println(s)
	/// This will print out “Hello”
	s, ok := i.(fileLogger)
	/// Type assertion also returns true/false depending on if the underlying type exists
	fmt.Println(s, ok)
	/// This will print out “Hello true”

	determineLogger(cl)
	determineLogger(fl)
}

func determineLogger(l logger) string {
	switch v := l.(type) {
	case fileLogger:
		return "It's a file logger!"
	case consoleLogger:
		return "It's a console logger!"
	default:
		fmt.Printf("Type %T! logger\n", v)
		return "It's an unknown logger!"
	}
}
