package helpers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/michaelwp/golang-gorm/models"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func EncryptPass(password []byte) ([]byte, error){
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {return nil, err}
	return hash, nil
}

func CompareHash(hash []byte, password []byte) error{
	err := bcrypt.CompareHashAndPassword(hash, password)
	if err != nil {return err}
	return nil
}

func CreateJwt(userId uint) (string, int64,  error) {
	SECRET_KEY := []byte(GetEnv("SECRET_KEY"))
	expirationTime := time.Now().Add(30 * time.Minute)

	claims := models.Claims{
		UserId: int(userId),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(SECRET_KEY)
	if err != nil {return "", 0, err}
	return tokenString, expirationTime.Unix(), nil
}

func SignedJwt(tokenString string) (uint, error){
	token, err := jwt.ParseWithClaims(
		tokenString,
		&models.Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(GetEnv("SECRET_KEY")), nil
		})

	var userId uint

	if err != nil {return 0, err}
	if claims, ok := token.Claims.(*models.Claims); ok && token.Valid {
		userId = uint(claims.UserId)
	}

	return userId, nil
}
