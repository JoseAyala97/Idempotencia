package middleware

import (
	"encoding/json"
	"log"
	"net/http"

	"IdEmpotencia/pkg/apperror"
)

func ErrorHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Println("Error interno del servidor:", err)
				respondWithAppError(w, apperror.InternalServerError())
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func respondWithAppError(w http.ResponseWriter, err *apperror.AppError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.Code)
	json.NewEncoder(w).Encode(map[string]string{"error": err.Message})
}
