package gin_helper

import (
	"github.com/gin-gonic/gin"
	"github.com/sudtanj/Go-SwissKnife/responder"
)

func recoveryHandler(c *gin.Context, err interface{}) {
	responder.Json(c, err)
}
