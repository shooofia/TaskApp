package model

type Task struct {
	ID             int    `json:"id" gorm:"primaryKey"`
	Judul          string `json:"judul"`
	Deskripsi      string `json:"deskripsi"`
	Prioritas      string `json:"prioritas"`
	TanggalTenggat string `json:"tanggal_tenggat"`
	Status         string `json:"status"`
}
