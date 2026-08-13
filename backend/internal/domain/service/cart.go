package service

import (
	"encoding/json"
	"fmt"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
	"github.com/nickznew1/MagazineMZM/backend/internal/domain/usecase"
	"net/http"
)

type CartService struct {
	useCase usecase.CartUseCase
}

func NewCartService(c *usecase.CartUseCase) *CartService {
	return &CartService{
		useCase: *c}
}

func (h *CartService) GetCart(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get shopping cart")
	userIdJWT := r.Context().Value("user_id")
	fmt.Println("userIdJWT for cart ", userIdJWT)
	idStr, ok := userIdJWT.(string)
	if !ok {
		RespondWithError(w, http.StatusForbidden, "invalid user id format")
		return
	}
	userId, err := h.useCase.GetCart(r.Context(), idStr)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "id doesnt find")
		return
	}
	RespondWithJSON(w, http.StatusOK, userId)
}

func (h *CartService) CreateUserItem(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CreateUser item")
	var input model.Cart
	fmt.Println(input)
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Неверные данные")
		return
	}
	shopCart, err := h.useCase.AddItemToCart(r.Context(), input)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "oshibka")
		return
	}
	RespondWithJSON(w, http.StatusCreated, shopCart)
}

func (h *CartService) DeleteUserItem(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DeleteUserItem")
	var input model.Cart

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Неверные данные")
		return
	}
	deleteItem, err := h.useCase.DeleteUserItem(r.Context(), input)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "error when deleting item from cart")
		return
	}
	RespondWithJSON(w, http.StatusNoContent, deleteItem)
}

func (h *CartService) CalcUserItem(w http.ResponseWriter, r *http.Request) {
	fmt.Println("calc item")
	var input model.Cart

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		RespondWithError(w, http.StatusBadRequest, "wrong body for counter")
		return
	}
	calcItem, err := h.useCase.CalcUserItem(r.Context(), input)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "error on server when try to count")
		return
	}
	RespondWithJSON(w, http.StatusCreated, calcItem)
}
