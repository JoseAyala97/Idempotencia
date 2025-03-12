package middleware

import (
	"IdEmpotencia/pkg/database"
	"context"
	"encoding/json"
	"net/http"
	"time"
)

type IdempotencyData struct {
	Status   string `json:"status"`
	Response string `json:"response"`
}

func IdempotencyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		idempotencyKey := r.Header.Get("Idempotency-Key")

		if idempotencyKey == "" {
			http.Error(w, "Idempotency-Key requerido", http.StatusBadRequest)
			return
		}

		data, err := database.RedisClient.Get(ctx, idempotencyKey).Result()
		if err == nil {
			var storedData IdempotencyData
			json.Unmarshal([]byte(data), &storedData)
			if storedData.Status == "IN_PROGRESS" {
				http.Error(w, "Solicitud en progreso", http.StatusConflict)
				return
			}
			w.Write([]byte(storedData.Response))
			return
		}

		inProgressData := IdempotencyData{Status: "IN_PROGRESS"}
		jsonData, _ := json.Marshal(inProgressData)
		database.RedisClient.Set(ctx, idempotencyKey, jsonData, 24*time.Hour)

		rec := &responseRecorder{ResponseWriter: w, statusCode: http.StatusOK}
		next.ServeHTTP(rec, r)

		finalData := IdempotencyData{
			Status:   "COMPLETED",
			Response: rec.responseBody,
		}
		finalJson, _ := json.Marshal(finalData)
		database.RedisClient.Set(ctx, idempotencyKey, finalJson, 24*time.Hour)
	})
}

type responseRecorder struct {
	http.ResponseWriter
	statusCode   int
	responseBody string
}

func (rr *responseRecorder) Write(b []byte) (int, error) {
	rr.responseBody = string(b)
	return rr.ResponseWriter.Write(b)
}

func (rr *responseRecorder) WriteHeader(statusCode int) {
	rr.statusCode = statusCode
	rr.ResponseWriter.WriteHeader(statusCode)
}
