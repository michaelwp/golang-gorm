package middlewares

import (
	"encoding/json"
	"github.com/michaelwp/golang-gorm/helpers"
	"github.com/michaelwp/golang-gorm/models"
	"net/http"
)

func IsAuthorized(next http.Handler) http.Handler {
	var res models.Result

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header["Token"]

		w.Header().Set("content-type", "application/json")
		res.Status = 0
		w.WriteHeader(http.StatusUnauthorized)

		if len(token) <= 0 {
			res.Message = "Token required"
			_ = json.NewEncoder(w).Encode(res)
			return
		} else {
			_, err := helpers.SignedJwt(token[0])
			if err != nil {
				res.Message = "Token Invalid"
				_ = json.NewEncoder(w).Encode(res)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}
