package repository

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	ID            uint      `gorm:"primarykey"`
	Nama          string    `json:"nama"`
	Tanggal_lahir time.Time `json:"tanggal_lahir"`
	Gender        int       `json:"gender"`
	ID_Hobi       int       `json:"id_hobi"`
	ID_Jurusan    int       `json:"id_jurusan"`
}

type Account struct {
	gorm.Model
	Fullname  string `gorm:"size:30"`
	Email     string `gorm:"size:50;primaryKey"`
	Password  string `gorm:"size:255"`
	Role      int    `gorm:"size:1"`
	Url_photo string `gorm:"size:255"`
}

type Role struct {
	ID   uint   `gorm:"primarykey"`
	Name string `gorm:"size:10"`
}
