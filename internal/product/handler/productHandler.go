package handler

import (
	"IdEmpotencia/internal/product/service"
	"IdEmpotencia/pkg/apperror"
	"IdEmpotencia/pkg/validate"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ProductHandler struct {
	service *service.ProductService
}

func NewProductHandler(service *service.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.service.GetAllProducts()
	if err != nil {
		respondWithAppError(w, apperror.InternalServerError())
		return
	}

	respondWithJSON(w, http.StatusOK, products)
}

func (h *ProductHandler) UpdateStock(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithAppError(w, apperror.BadRequest("ID inválido"))
		return
	}

	var req struct {
		Stock int `json:"stock" validate:"required,gte=0"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithAppError(w, apperror.BadRequest("Error en el JSON"))
		return
	}

	if msg, err := validate.Validate(req); err != nil {
		respondWithAppError(w, apperror.BadRequest(msg))
		return
	}

	if err := h.service.UpdateStock(id, req.Stock); err != nil {
		respondWithAppError(w, apperror.InternalServerError())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Stock actualizado"})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}

func respondWithAppError(w http.ResponseWriter, err *apperror.AppError) {
	respondWithJSON(w, err.Code, map[string]string{"error": err.Message})
}
