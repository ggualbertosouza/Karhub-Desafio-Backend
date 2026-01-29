package HttpContext

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func responseError(
	c *gin.Context,
	status int,
	message string,
) {
	c.AbortWithStatusJSON(status, gin.H{
		"error": message,
	})
}

func BadRequest(c *gin.Context, message string) {
	responseError(c, http.StatusBadRequest, message)
}

func NotFound(c *gin.Context, message string) {
	responseError(c, http.StatusNotFound, message)
}

func Conlfict(c *gin.Context, message string) {
	responseError(c, http.StatusConflict, message)
}
