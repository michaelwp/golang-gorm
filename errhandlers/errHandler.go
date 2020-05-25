package errhandlers

import (
	"encoding/json"
	"github.com/michaelwp/golang-gorm/models"
	"log"
	"net/http"
)

func ErrCreate(w http.ResponseWriter, res models.Result, msg string) {
	res.Status = 0
	res.Message = msg
	w.WriteHeader(http.StatusBadRequest)
	err := json.NewEncoder(w).Encode(res)
	if err != nil {log.Fatal(err)}
}
