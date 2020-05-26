package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/michaelwp/golang-gorm/db"
	"github.com/michaelwp/golang-gorm/errhandlers"
	"github.com/michaelwp/golang-gorm/helpers"
	"github.com/michaelwp/golang-gorm/models"
	"log"
	"net/http"
	"strings"
)

func Home(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Home")
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	var u models.UserInput
	var res models.Result

	res.Status = 1
	res.Message = "Data successfully inserted"
	w.Header().Set("Content-type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {log.Fatal(err)}

	_, err = findEmail(u.Email)

	if err == nil {
		errhandlers.ErrCreate(w, res, "Email already registered")
		return
	}

	user := models.User{
		FirstName: strings.ToLower(u.FirstName),
		LastName: strings.ToLower(u.LastName),
	}

	userRes := db.MySql().Create(&user)
	if userRes.Error != nil {
		errhandlers.ErrCreate(w, res, fmt.Sprintf("%s", userRes.Error))
		return
	} else {

		hash, err := helpers.EncryptPass([]byte(u.Password))
		if err != nil {log.Fatal(err)}

		cred := models.Credential{
			UserID: user.ID,
			Email: strings.ToLower(u.Email),
			Password: string(hash),
		}

		credRes := db.MySql().Create(&cred)
		if credRes.Error != nil {
			errhandlers.ErrCreate(w, res, fmt.Sprintf("%s", credRes.Error))
			return
		}
	}

	defer db.MySql().Close()

	err = json.NewEncoder(w).Encode(res)
	if err != nil {log.Fatal(err)}
}

func findEmail(email string) (models.Credential, error) {
	var c models.Credential

	credRes := db.MySql().
		Where("email = ?", email).
		Find(&c)

	return c, credRes.Error
}

func Login(w http.ResponseWriter, r *http.Request) {
	var c models.Credential
	var res models.ResultToken
	var tokenData models.TokenData
	errMsg := "Email/password not found"

	res.Status = 1
	res.Message = "successfully login"

	w.Header().Set("content-type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {log.Fatal(err)}

	cred, err := findEmail(c.Email)
	if err != nil {
		res.Status = 0
		res.Message = errMsg
		w.WriteHeader(http.StatusUnauthorized)
	}

	err = helpers.CompareHash([]byte(cred.Password), []byte(c.Password))
	if err != nil {
		res.Status = 0
		res.Message = errMsg
		w.WriteHeader(http.StatusUnauthorized)
	} else {
		token, exp, err := helpers.CreateJwt(c.UserID)
		if err != nil {log.Println(err)}

		tokenData.ExpiresAt = exp
		tokenData.Token = token
		res.Data = tokenData
	}

	_ = json.NewEncoder(w).Encode(res)
}