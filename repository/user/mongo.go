package user

import (
	"api-desatanggap/business/user"
	"api-desatanggap/repository"
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
	colCode *mongo.Collection
}

func NewMongoRepository(col *mongo.Database) *MongoDBRepository {
	return &MongoDBRepository{
		col:     col.Collection("users"),
		colRole: col.Collection("roles_user"),
		colCode: col.Collection("code_otp"),
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
	if len(tmpAccount) == 0 {
		return nil, errors.New("Data Not Found")
	}
	return &tmpAccount[0], nil
}

func (repo *MongoDBRepository) CreateAccount(Data *user.RegAccount) (*user.Account, error) {
	hashpw, _ := utils.Hash(Data.Password)
	ObjId_userid, err := primitive.ObjectIDFromHex(Data.Role_id)
	if err != nil {
		return nil, err
	}
	InsertData := &repository.Account{
		Email:      Data.Email,
		Fullname:   Data.Fullname,
		Password:   string(hashpw),
		Role_id:    ObjId_userid,
		IsVerified: false,
	}
	result, err := repo.col.InsertOne(context.Background(), InsertData)
	if err != nil {
		return nil, err
	}
	id, err := primitive.ObjectIDFromHex(fmt.Sprintf("%s", result.InsertedID))

	InsertData.ID = id

	ResponseAccount := &user.Account{
		ID:         id,
		Email:      InsertData.Email,
		Fullname:   InsertData.Fullname,
		Password:   InsertData.Password,
		Role_id:    InsertData.Role_id,
		IsVerified: InsertData.IsVerified,
	}
	return ResponseAccount, nil
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

func (repo *MongoDBRepository) SendVerification(email string) error {
	var tmpAcc repository.Account
	err := repo.col.FindOne(context.Background(), bson.M{"email": email}).Decode(&tmpAcc)
	if err != nil {
		return errors.New("Email Not Registered")
	}
	codeotp, err := utils.InitEmail(email)
	if err != nil {
		return err
	}
	err = repo.CreateCodeOtp(email, codeotp)
	return nil
}

func (repo *MongoDBRepository) ValidationEmail(Data string) error {
	return nil
}

func (repo *MongoDBRepository) CreateCodeOtp(email string, codeotp string) error {
	timeExpired := time.Now().Add(24 * time.Hour)
	InsertCode := &repository.CodeOtp{
		Email:      email,
		Code:       codeotp,
		Expired_at: timeExpired,
	}
	var tmpAcc repository.CodeOtp
	repo.colCode.FindOne(context.Background(), bson.M{"email": email}).Decode(&tmpAcc)
	if tmpAcc.Email != "" {
		filter := bson.M{"email": email}
		update := bson.M{
			"$set": bson.M{
				"code":       codeotp,
				"expired_at": timeExpired,
			},
		}
		repo.colCode.UpdateOne(context.Background(), filter, update)
		return nil
	}
	_, err := repo.colCode.InsertOne(context.Background(), InsertCode)
	if err != nil {
		return err
	}
	return nil
}

func (repo *MongoDBRepository) VerificationAccount(code string) error {
	var codeotp repository.CodeOtp
	err := repo.colCode.FindOneAndDelete(context.Background(), bson.M{"code": code}).Decode(&codeotp)
	if err != nil {
		return errors.New("Code Not Found")
	}
	if codeotp.Expired_at.Before(time.Now()) {
		return errors.New("Code Expired")
	}
	filter := bson.M{"email": codeotp.Email}
	update := bson.M{"isverified": true}
	_, err = repo.col.UpdateOne(context.Background(), filter, bson.M{"$set": update})
	if err != nil {
		return err
	}
	return nil
}

func (repo *MongoDBRepository) DeleteUser(email string) error {
	filter := bson.M{"email": email}
	err := repo.col.FindOneAndDelete(context.Background(), filter).Err()
	if err != nil {
		return errors.New("Data Not Found")
	}
	return nil
}
