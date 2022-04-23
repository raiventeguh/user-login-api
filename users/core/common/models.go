package common

import (
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id       primitive.ObjectID `json:"id,omitempty"`
	Name     string             `json:"name,omitempty" validate:"required"`
	Email    string             `json:"email,omitempty" validate:"required"`
	Password string             `json:"password,omitempty" validate:"required"`
	Age      *int               `json:"age,omitempty" validate:"required"`
	Admin    bool               `json:"admin,omitempty"`
}

type Auth struct {
	Email    string `json:"email,omitempty"" validate:"required"`
	Password string `json:"password,omitempty" validate:"required"`
}

type JwtCustomClaims struct {
	Name   string `json:"name"`
	Admin  bool   `json:"admin"`
	UserId string `json:"id"`
	jwt.StandardClaims
}
