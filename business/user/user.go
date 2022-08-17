package user

import (
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Customer struct {
	ID            uint   `json:"id"`
	Nama          string `json:"nama"`
	Tanggal_lahir string `json:"tanggal_lahir"`
	Gender        string `json:"gender"`
	ID_Hobi       int    `json:"id_hobi"`
	ID_Gender     int    `json:"id_gender"`
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
	Role_id   string `json:"role_id,omitempty" binding:"required"`
	Url_photo string `json:"url_photo"`
}

type Account struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Email    string             `bson:"email,omitempty" binding:"required"`
	Fullname string             `bson:"fullname,omitempty" binding:"required"`
	Password string             `bson:"password,omitempty" binding:"required"`
	Role_id  primitive.ObjectID `bson:"role_id,omitempty" binding:"required"`
	Roles    []Role             `bson:"roles" json:"roles"`
}

type Role struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Name        string             `bson:"name,omitempty" binding:"required" json:"rolename"`
	Label       string             `bson:"label,omitempty" binding:"required" json:"rolelabel"`
	Description string             `bson:"description,omitempty" binding:"required" json:"description"`
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
