package controllers

import (
	"encoding/json"
	"fmt"
	"golang-gorm/db"
	"golang-gorm/errhandlers"
	"golang-gorm/models"
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

	mySql := db.MySql()

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {log.Fatal(err)}

	user := models.User{
		FirstName: strings.ToLower(u.FirstName),
		LastName: strings.ToLower(u.LastName),
	}

	userRes := mySql.Create(&user)
	if userRes.Error != nil {
		errhandlers.ErrCreate(w, res, fmt.Sprintf("%s", userRes.Error))
		return
	} else {
		cred := models.Credential{
			UserID: user.ID,
			Email: strings.ToLower(u.Email),
			Password: u.Password,
		}

		credRes := mySql.Create(&cred)
		if credRes.Error != nil {
			errhandlers.ErrCreate(w, res, fmt.Sprintf("%s", credRes.Error))
			return
		}
	}

	mySql.Close()

	err = json.NewEncoder(w).Encode(res)
	if err != nil {log.Fatal(err)}
}