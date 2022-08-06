package repository

import (
	"time"
)

type Customer struct {
	ID            uint      `gorm:"primarykey"`
	Nama          string    `json:"nama"`
	Tanggal_lahir time.Time `json:"tanggal_lahir"`
	Gender        int       `json:"gender"`
	ID_Hobi       int       `json:"id_hobi"`
	Hobi          Hobi      `gorm:"foreignkey:ID;references:ID_Hobi"`
	ID_Jurusan    int       `json:"id_jurusan"`
	Jurusan       Jurusan   `gorm:"foreignkey:ID;references:ID_Jurusan"`
}

type Jurusan struct {
	ID   int    `gorm:"primarykey"`
	Nama string `json:"nama"`
}

type Hobi struct {
	ID   int    `gorm:"primarykey"`
	Nama string `json:"nama"`
}
