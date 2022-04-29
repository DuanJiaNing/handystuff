package main

import (
	"fmt"
	"handystuff/api"
	"handystuff/api/middleware"
	"handystuff/common/log"
	"handystuff/common/statsd"
	"handystuff/config"
	"handystuff/stuff/decrypt"
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func setup(conf *config.Conf) (*gin.Engine, error) {
	engine := gin.Default()
	engine.Use(gin.ErrorLogger())

	logWriter := setupLogWriter()
	var err error
	var statsdClient api.StatsdClient
	statsdClient = statsd.NewDummyClient()
	if conf.Tick.Enable {
		statsdClient, err = statsd.NewClient(&conf.Tick)
		if err != nil {
			return nil, err
		}
	}

	mdw := middleware.NewMiddleward(statsdClient)
	engine.Use(mdw.Cost())

	handler := api.NewHandler(
		log.NewLogger(logWriter),
		statsdClient,
		decrypt.NewStuff(conf.AESCrypt),
	)
	setupRouter(engine, handler)

	return engine, nil
}

func setupRouter(engine *gin.Engine, handler api.Handler) {
	engine.GET("/", func(c *gin.Context) { c.File("cmd/api/static/index.html") })
	engine.POST("/decrypt/:app", handler.Decrypt)
	engine.POST("/encrypt/:app", handler.Encrypt)
	engine.GET("/tools/token", handler.GenerateToken)
}

func setupLogWriter() io.Writer {
	var logWriter []io.Writer
	logWriter = append(logWriter, os.Stdout)
	name := fmt.Sprintf("logs/app.%s.log", time.Now().Format("20060102"))
	f, _ := os.OpenFile(name, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0666)
	if f != nil {
		logWriter = append(logWriter, f)
	}
	gin.DefaultWriter = io.MultiWriter(logWriter...)
	gin.ForceConsoleColor()

	return gin.DefaultWriter
}
