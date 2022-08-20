package repository

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// type Customer struct {
// 	ID            uint      `gorm:"primarykey"`
// 	Nama          string    `json:"nama"`
// 	Tanggal_lahir time.Time `json:"tanggal_lahir"`
// 	Gender        int       `json:"gender"`
// 	ID_Hobi       int       `json:"id_hobi"`
// 	ID_Jurusan    int       `json:"id_jurusan"`
// }

type Account struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Email      string             `bson:"email,omitempty" binding:"required"`
	Fullname   string             `bson:"fullname,omitempty" binding:"required"`
	Password   string             `bson:"password,omitempty" binding:"required"`
	Role_id    primitive.ObjectID `bson:"role_id,omitempty" binding:"required"`
	IsVerified bool               `bson:"isverified"`
}

type Admin struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username,omitempty" binding:"required"`
	Fullname string             `bson:"fullname,omitempty" binding:"required"`
	Password string             `bson:"password,omitempty" binding:"required"`
	Role_id  primitive.ObjectID `bson:"role_id,omitempty" binding:"required"`
	// Roles     Role               `bson:"roles"`
}
type Role struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Rolename    string             `bson:"rolename,omitempty" binding:"required" json:"rolename"`
	Rolelabel   string             `bson:"rolelabel,omitempty" binding:"required" json:"rolelabel"`
	Description string             `bson:"description,omitempty" binding:"required" json:"description"`
}

type CodeOtp struct {
	ID         string    `json:"id" bson:"_id,omitempty"`
	Email      string    `json:"email" bson:"email,omitempty"`
	Code       string    `json:"code" bson:"code,omitempty"`
	Expired_at time.Time `json:"expired_at" bson:"expired_at,omitempty"`
}
