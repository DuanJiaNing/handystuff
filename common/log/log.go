package log

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Printlnf(format string, a ...interface{}) {
	str := fmt.Sprintln(fmt.Sprintf(format, a...))
	_, err := gin.DefaultWriter.Write([]byte(str))
	if err != nil {
		panic(err)
	}
}
