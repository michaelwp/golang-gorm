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

	err, _ = findEmail(u.Email)

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

func findEmail(email string) (error, models.Credential) {
	var c models.Credential

	credRes := db.MySql().
		Where("email like ?", "%" + email + "%").
		Find(&c)

	return credRes.Error, c
}