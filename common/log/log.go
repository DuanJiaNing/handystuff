package log

import (
	"fmt"
	"io"
)

type Logger struct {
	Writer io.Writer
}

func NewLogger(writer io.Writer) Logger {
	return Logger{
		Writer: writer,
	}
}

func (a Logger) Infof(format string, p ...interface{}) {
	str := fmt.Sprintln(fmt.Sprintf(format, p...))
	str = fmt.Sprintf("[%s]: %s", "INFO", str)
	a.log(str)
}

func (a Logger) Warnf(format string, p ...interface{}) {
	str := fmt.Sprintln(fmt.Sprintf(format, p...))
	str = fmt.Sprintf("[%s]: %s", "WARN", str)
	a.log(str)
}

func (a Logger) Debugf(format string, p ...interface{}) {
	str := fmt.Sprintln(fmt.Sprintf(format, p...))
	str = fmt.Sprintf("[%s]: %s", "DEBUG", str)
	a.log(str)
}

func (a Logger) log(message string) {
	_, err := a.Writer.Write([]byte(message))
	if err != nil {
		panic(err)
	}
}
