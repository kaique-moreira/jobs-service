package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	job := r.Group("/jobs")

	job.GET("/opening", func(ctx *gin.Context) {
		resp := "kaique"
		ctx.JSON(http.StatusOK, resp)
	})

	job.GET("/openings", func(ctx *gin.Context) {
		resp := "kaique"
		ctx.JSON(http.StatusOK, resp)
	})

	job.POST("/opening", func(ctx *gin.Context) {
		resp := "kaique"
		ctx.JSON(http.StatusOK, resp)
	})

	job.DELETE("/opening", func(ctx *gin.Context) {
		resp := "kaique"
		ctx.JSON(http.StatusOK, resp)
	})

	job.PUT("/opening", func(ctx *gin.Context) {
		resp := "kaique"
		ctx.JSON(http.StatusOK, resp)
	})
}
