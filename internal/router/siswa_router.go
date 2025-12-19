package router

import (
	"github.com/entertrans/go-base-project.git/internal/handler"
	"github.com/entertrans/go-base-project.git/internal/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterSiswaRoutes(rg *gin.RouterGroup, siswaHandler handler.SiswaHandler, jwtSecret string) {
	siswa := rg.Group("/siswa")
	siswa.Use(middleware.AuthMiddleware(jwtSecret))

	{
		// Get all siswa
		siswa.GET("/", siswaHandler.GetAllSiswa)

		// Nanti bisa ditambahkan endpoint lain di sini
		// siswa.GET("/aktif", siswaHandler.GetSiswaAktif)
		// siswa.GET("/:nis", siswaHandler.GetSiswaByNIS)
	}
}