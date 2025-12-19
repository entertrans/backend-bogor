package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/entertrans/go-base-project.git/internal/controller"
)

type SiswaHandler interface {
	GetAllSiswa(c *gin.Context)
}

type siswaHandler struct {
	siswaController controller.SiswaController
}

func NewSiswaHandler(siswaController controller.SiswaController) SiswaHandler {
	return &siswaHandler{siswaController: siswaController}
}

// GetAllSiswa handler untuk mendapatkan semua siswa
func (h *siswaHandler) GetAllSiswa(c *gin.Context) {
	siswaList, err := h.siswaController.GetAllSiswa()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Data siswa berhasil diambil",
		"data":    siswaList,
	})
}