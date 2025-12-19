package controller

import (
	"gorm.io/gorm"

	"github.com/entertrans/go-base-project.git/internal/dto"
	"github.com/entertrans/go-base-project.git/internal/model"
)

type SiswaController interface {
	GetAllSiswa() ([]dto.SiswaResponse, error)
}

type siswaController struct {
	db *gorm.DB
}

func NewSiswaController(db *gorm.DB) SiswaController {
	return &siswaController{
		db: db,
	}
}

// GetAllSiswa controller untuk mendapatkan semua siswa
func (c *siswaController) GetAllSiswa() ([]dto.SiswaResponse, error) {
	var siswaList []model.Siswa
	
	// Query dengan preload kelas
	err := c.db.Preload("Kelas").
		Find(&siswaList).Error
	
	if err != nil {
		return nil, err
	}
	
	// Convert ke response DTO
	var responses []dto.SiswaResponse
	for _, siswa := range siswaList {
		responses = append(responses, dto.SiswaResponse{
			SiswaID:      siswa.SiswaID,
			SiswaNIS:     siswa.SiswaNIS,
			SiswaNama:    siswa.SiswaNama,
			SiswaJenkel:  siswa.SiswaJenkel,
			SiswaKelasID: siswa.SiswaKelasID,
			SoftDeleted:  siswa.SoftDeleted,
			Kelas:        siswa.Kelas,
		})
	}
	
	return responses, nil
}