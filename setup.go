package main

import (
	"fmt"
	"handystuff/common/statsd"
	"handystuff/config"
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func setup(configFile string) error {
	var err error

	if err = config.Load(configFile); err != nil {
		return err
	}
	setupLogWriter()
	if err = statsd.Init(); err != nil {
		return err
	}

	return nil
}

func setupLogWriter() {
	var logWriter []io.Writer
	logWriter = append(logWriter, os.Stdout)
	name := fmt.Sprintf("logs/app.log.%s", time.Now().Format("20060102"))
	f, _ := os.OpenFile(name, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0666)
	if f != nil {
		logWriter = append(logWriter, f)
	}
	gin.DefaultWriter = io.MultiWriter(logWriter...)
	gin.ForceConsoleColor()
}
