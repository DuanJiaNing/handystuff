package main

import (
	"fmt"
	"handystuff/config"
	"handystuff/stuff/decrypt"
	"handystuff/stuff/middleware"
	"handystuff/stuff/tools"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
)

func main() {
	configFile := pflag.StringP("config", "c", "config.dev.yaml", "specific config file.")
	err := setup(*configFile)
	if err != nil {
		panic(err)
	}

	router := gin.Default()
	router.Use(gin.ErrorLogger())
	router.Use(middleware.Cost())
	router.GET("/", func(c *gin.Context) { c.File("static/index.html") })
	router.POST("/decrypt/:app", decrypt.Decrypt)
	router.POST("/encrypt/:app", decrypt.Encrypt)
	router.GET("/tools/token", tools.GenerateToken)

	log.Fatalln(router.Run(fmt.Sprintf(":%d", config.Conf.App.Port)))
}
