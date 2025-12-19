package dto

import "github.com/entertrans/go-base-project.git/internal/model"

type SiswaResponse struct {
	SiswaID      uint    `json:"siswa_id"`
	SiswaNIS     *string `json:"siswa_nis"`
	SiswaNama    *string `json:"siswa_nama"`
	SiswaJenkel  *string `json:"siswa_jenkel"`
	SiswaKelasID *int    `json:"siswa_kelas_id"`
	SoftDeleted  *int    `json:"soft_deleted"`
	Kelas        model.Kelas `json:"kelas"`
}