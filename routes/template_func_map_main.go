package routes

import (
	"text/template"

	"github.com/gin-gonic/gin"
)

func SetFuncTemplate(router *gin.Engine) {
	router.SetFuncMap(template.FuncMap{})
}
