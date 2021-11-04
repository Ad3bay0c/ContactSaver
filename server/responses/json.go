package responses

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func JSON(c *gin.Context,code int, message interface{}, err string, data interface{}) {
	response := gin.H{
		"Error": err,
		"message": message,
		"data": data,
		"status": http.StatusText(code),
		"time_stamp": time.Now().Format("2006-01-02 15:04:05"),
	}
	c.JSON(code, response)
}