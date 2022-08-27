package admin

import (
	"time"

	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthLogin struct {
	Username string `bson:"username,omitempty" validate:"required"`
	Password string `bson:"password,omitempty" validate:"required"`
}

type Admin struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username,omitempty" binding:"required"`
	Fullname string             `bson:"fullname,omitempty" binding:"required"`
	Password string             `bson:"password,omitempty" binding:"required"`
	Role_id  primitive.ObjectID `bson:"role_id,omitempty" binding:"required" json:"role_id"`
	Roles    []Role             `bson:"roles" json:"roles"`
}

type Role struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Name        string             `bson:"name,omitempty" binding:"required" json:"name"`
	Label       string             `bson:"label,omitempty" binding:"required" json:"label"`
	Description string             `bson:"description,omitempty" binding:"required" json:"description"`
}

type RegAdmin struct {
	Username string `bson:"username,omitempty" binding:"required"`
	Fullname string `bson:"fullname,omitempty" binding:"required"`
	Password string `bson:"password,omitempty" binding:"required"`
	Role_id  string `bson:"role_id,omitempty" binding:"required"`
}

type ResponseLogin struct {
	Admin Admin  `json:"admin"`
	Token string `json:"token"`
}

type Claims struct {
	Username string
	Email    string
	Role     string
	jwt.StandardClaims
}

type Cooperation struct {
	ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name       string             `json:"name" bson:"name,omitempty"`
	Address    string             `json:"address" bson:"address,omitempty"`
	City       string             `json:"city" bson:"city,omitempty"`
	Province   string             `json:"province" bson:"province,omitempty"`
	Latitude   string             `json:"latitude" bson:"latitude,omitempty"`
	Longitude  string             `json:"longitude" bson:"longitude,omitempty"`
	Photo_url  string             `json:"photo_url" bson:"photo_url,omitempty"`
	Email      string             `json:"email" bson:"email,omitempty"`
	Username   string             `json:"username" bson:"username,omitempty"`
	Password   string             `json:"password" bson:"password,omitempty"`
	Created_at time.Time          `json:"created_at" bson:"created_at,omitempty"`
}

type RegCooperation struct {
	Name       string             `json:"name" bson:"name,omitempty"`
	Address    string             `json:"address" bson:"address,omitempty"`
	City       string             `json:"city" bson:"city,omitempty"`
	Province   string             `json:"province" bson:"province,omitempty"`
	Latitude   string             `json:"latitude" bson:"latitude,omitempty"`
	Longitude  string             `json:"longitude" bson:"longitude,omitempty"`
	Photo_url  string             `json:"photo_url" bson:"photo_url,omitempty"`
	Email      string             `json:"email" bson:"email,omitempty"`
	Username   string             `json:"username" bson:"username,omitempty"`
	Password   string             `json:"password" bson:"password,omitempty"`
	Role_id    primitive.ObjectID `bson:"role_id,omitempty" binding:"required" json:"role_id"`
	Created_at time.Time          `json:"created_at" bson:"created_at,omitempty"`
}
