package logic

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(context *gin.Context) {
	context.HTML(http.StatusOK, "index.tmpl", nil)
}
