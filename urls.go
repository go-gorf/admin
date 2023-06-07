package admin

import (
	"github.com/gin-gonic/gin"
)

func Urls(r *gin.Engine) {
	r.Static("/admin", StaticFilesPath())
}
