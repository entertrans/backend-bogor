package router

import (
	"github.com/gin-gonic/gin"

	"github.com/entertrans/go-base-project.git/internal/handler"
)

func RegisterPingRoutes(rg *gin.RouterGroup, pingHandler handler.PingHandler) {
	rg.GET("/ping", pingHandler.Ping)
	
	// Jika Anda ingin menambahkan HealthCheck, Version, dll
	// Anda perlu menambahkan method tersebut ke interface PingHandler terlebih dahulu
}