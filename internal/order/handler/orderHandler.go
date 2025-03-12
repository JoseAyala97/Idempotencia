package handler

import (
	"IdEmpotencia/internal/order"
	"IdEmpotencia/internal/order/service"
	"IdEmpotencia/pkg/apperror"
	"IdEmpotencia/pkg/validate"
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type OrderHandler struct {
	service *service.OrderService
}

func NewOrderHandler(service *service.OrderService) *OrderHandler {
	return &OrderHandler{service: service}
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var o order.Order
	if err := json.NewDecoder(r.Body).Decode(&o); err != nil {
		respondWithAppError(w, apperror.BadRequest("Error en el JSON"))
		return
	}

	if msg, err := validate.Validate(o); err != nil {
		respondWithAppError(w, apperror.BadRequest(msg))
		return
	}

	idempotencyKey := r.Header.Get("Idempotency-Key")
	if idempotencyKey == "" {
		respondWithAppError(w, apperror.BadRequest("Idempotency-Key requerido"))
		return
	}

	orderID, err := h.service.CreateOrder(context.Background(), idempotencyKey, &o)
	if err != nil {
		respondWithAppError(w, apperror.NewAppError(http.StatusConflict, err.Error()))
		return
	}

	respondWithJSON(w, http.StatusCreated, map[string]int{"order_id": orderID})
}

func (h *OrderHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithAppError(w, apperror.BadRequest("ID inválido"))
		return
	}

	o, err := h.service.GetOrderByID(id)
	if err != nil {
		respondWithAppError(w, apperror.NotFound("Pedido no encontrado"))
		return
	}

	respondWithJSON(w, http.StatusOK, o)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}

func respondWithAppError(w http.ResponseWriter, err *apperror.AppError) {
	respondWithJSON(w, err.Code, map[string]string{"error": err.Message})
}
