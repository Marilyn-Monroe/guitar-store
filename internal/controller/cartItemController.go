package controller

import (
	"encoding/json"
	"fmt"
	"guitarStore/api"
	"guitarStore/internal/common"
	"guitarStore/internal/model"
	"guitarStore/internal/service"
	"net/http"
)

type CartItemController struct {
	service.CartItemService
}

func NewCartItemController(cartItemService *service.CartItemService) *CartItemController {
	return &CartItemController{CartItemService: *cartItemService}
}

func (cartItemController CartItemController) GetCartItemsByUserId(w http.ResponseWriter, r *http.Request, params api.GetCartItemsByUserIdParams) {
	ctx := r.Context()

	cartItems, err := cartItemController.CartItemService.FindByUserId(ctx, params.UserId)
	if err != nil {
		common.SendResponse(w, http.StatusInternalServerError, "Failed to retrieve cart items")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Pagination-Count", fmt.Sprint(0))
	w.Header().Set("X-Pagination-Limit", fmt.Sprint(0))
	w.Header().Set("X-Pagination-Page", fmt.Sprint(0))
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(cartItems); err != nil {
		common.SendResponse(w, http.StatusInternalServerError, "Failed to encode response")
	}
}

func (cartItemController CartItemController) EditCartItem(w http.ResponseWriter, r *http.Request) {
	var requestBody model.CartItemModel
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		common.SendResponse(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	updatedCartItem, err := cartItemController.Update(r.Context(), requestBody)
	if err != nil {
		common.SendResponse(w, http.StatusInternalServerError, "Failed to update cart item")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(updatedCartItem); err != nil {
		common.SendResponse(w, http.StatusInternalServerError, "Failed to encode response")
	}
}
