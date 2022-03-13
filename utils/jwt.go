package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/karim-w/techunicorn-task/models"
)

func CreateToken(userModel *models.User) (string, error) {
	var err error
	//Creating Access Token
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd") //this should be in an env file
	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = userModel.ID
	atClaims["exp"] = time.Now().Add(time.Hour * 2).Unix()
	atClaims["iat"] = time.Now().Unix()
	atClaims["type"] = userModel.Type
	atClaims["nbf"] = time.Now().Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}
