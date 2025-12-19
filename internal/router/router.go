package router

import (
	"github.com/gin-gonic/gin"

	"github.com/entertrans/go-base-project.git/internal/handler"

)

type Router struct {
	pingHandler  handler.PingHandler
	authHandler  handler.AuthHandler
	siswaHandler handler.SiswaHandler
	jwtSecret    string
}

func NewRouter(
	pingHandler handler.PingHandler, 
	authHandler handler.AuthHandler, 
	siswaHandler handler.SiswaHandler,
	jwtSecret string,
) *Router {
	return &Router{
		pingHandler:  pingHandler,
		authHandler:  authHandler,
		siswaHandler: siswaHandler,
		jwtSecret:    jwtSecret,
	}
}

func (r *Router) SetupRoutes(engine *gin.Engine) {
	api := engine.Group("/api/v1")
	
	// Register all route groups
	RegisterPingRoutes(api, r.pingHandler)
	RegisterAuthRoutes(api, r.authHandler, r.jwtSecret)
	RegisterSiswaRoutes(api, r.siswaHandler, r.jwtSecret)
}

// Tambahkan deklarasi fungsi-fungsi register yang digunakan
// func RegisterSiswaRoutes(rg *gin.RouterGroup, siswaHandler handler.SiswaHandler, jwtSecret string)