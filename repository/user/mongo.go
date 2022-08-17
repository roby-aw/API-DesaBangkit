package user

import (
	"api-desatanggap/business/user"
	"api-desatanggap/utils"
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBRepository struct {
	col     *mongo.Collection
	colRole *mongo.Collection
}

func NewMongoRepository(col *mongo.Database) *MongoDBRepository {
	return &MongoDBRepository{
		col:     col.Collection("users"),
		colRole: col.Collection("roles_user"),
	}
}

func (repo *MongoDBRepository) FindAccountByEmail(email string) (*user.Account, error) {
	filter := bson.A{
		bson.M{
			"$match": bson.M{
				"email": bson.M{"$regex": email},
			},
		},
		bson.M{
			"$lookup": bson.M{
				"from":         "roles_user",
				"localField":   "role_id",
				"foreignField": "_id",
				"as":           "roles",
			},
		},
	}
	cur, err := repo.col.Aggregate(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	var tmpAccount []user.Account
	if err = cur.All(context.Background(), &tmpAccount); err != nil {
		return nil, err
	}
	if len(tmpAccount) < 1 {
		return nil, errors.New("Data Not Found")
	}
	return &tmpAccount[0], nil
}

func (repo *MongoDBRepository) CreateAccount(Data *user.RegAccount) (*user.Account, error) {
	hashpw, _ := utils.Hash(Data.Password)
	ObjId_userid, _ := primitive.ObjectIDFromHex(Data.Role_id)
	InsertData := &user.Account{
		Email:    Data.Email,
		Fullname: Data.Fullname,
		Password: string(hashpw),
		Role_id:  ObjId_userid,
	}
	result, err := repo.col.InsertOne(context.Background(), InsertData)
	if err != nil {
		return nil, err
	}
	id, err := primitive.ObjectIDFromHex(fmt.Sprintf("%s", result.InsertedID))

	InsertData.ID = id
	return InsertData, nil
}

func (repo *MongoDBRepository) CreateToken(Data *user.Account) (*string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &user.Claims{
		Email: Data.Email,
		Role:  Data.Roles[0].Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	SECRET_KEY := os.Getenv("SECRET_JWT")
	token_jwt, err := token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		return nil, err
	}
	return &token_jwt, err
}

func (repo *MongoDBRepository) Createcustomer(Data *user.Regcustomer) (*user.Regcustomer, error) {
	return nil, nil
}

func (repo *MongoDBRepository) Findcustomer() ([]user.Customer, error) {
	return nil, nil
}

func (repo *MongoDBRepository) GetRole() ([]*user.Role, error) {
	var Role []*user.Role
	cur, err := repo.colRole.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	cur.All(context.Background(), &Role)
	return Role, err
}
