package customer

import "github.com/golang-jwt/jwt/v4"

type Customer struct {
	ID            uint   `json:"id"`
	Nama          string `json:"nama"`
	Tanggal_lahir string `json:"tanggal_lahir"`
	Gender        string `json:"gender"`
	ID_Hobi       int    `json:"id_hobi"`
	ID_Gender     int    `json:"id_gender"`
}

type Detail_customer struct {
	ID            uint    `gorm:"primarykey"`
	Nama          string  `json:"nama"`
	Tanggal_lahir string  `json:"tanggal_lahir"`
	Gender        string  `json:"gender"`
	ID_Hobi       int     `json:"id_hobi"`
	Hobi          Hobi    `gorm:"foreignkey:ID;references:ID_Hobi"`
	ID_Jurusan    int     `json:"id_jurusan"`
	Jurusan       Jurusan `gorm:"foreignkey:ID;references:ID_Jurusan"`
}

type Regcustomer struct {
	Nama          string `json:"nama"`
	Tanggal_lahir string `json:"tanggal_lahir"`
	Gender        int    `json:"gender"`
}

type RegAccount struct {
	Fullname  string `json:"fullname" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
	ID_Role   int    `json:"role" validate:"required"`
	Url_photo string `json:"url_photo"`
}

type Account struct {
	ID        uint   `json:"id"`
	Fullname  string `json:"fullname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	ID_Role   int    `json:"id_role"`
	Role      Role   `json:"role" gorm:"foreignkey:ID;references:ID_Role"`
	Url_photo string `json:"url_photo"`
}

type Role struct {
	ID   uint   `gorm:"primarykey"`
	Name string `gorm:"size:10"`
}

type AuthLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type Claims struct {
	ID    int
	Email string
	Role  string
	jwt.StandardClaims
}

type ResLogin struct {
	Account Account `json:"account"`
	Token   string  `json:"token"`
}

type Jurusan struct {
	ID   string `json:"id"`
	Nama string `json:"nama"`
}

type Hobi struct {
	ID   int    `gorm:"primarykey" json:"id"`
	Nama string `json:"nama"`
}
