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

type ReviewController struct {
	service.ReviewService
}

func NewReviewController(reviewService *service.ReviewService) *ReviewController {
	return &ReviewController{ReviewService: *reviewService}
}

func (reviewController ReviewController) FindReviewsByGuitarId(w http.ResponseWriter, r *http.Request, params api.FindReviewsByGuitarIdParams) {
	ctx := r.Context()

	reviews, err := reviewController.ReviewService.FindByGuitarId(ctx, params.GuitarId)
	if err != nil {
		common.SendResponse(w, http.StatusInternalServerError, "Failed to retrieve reviews")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Pagination-Count", fmt.Sprint(0))
	w.Header().Set("X-Pagination-Limit", fmt.Sprint(0))
	w.Header().Set("X-Pagination-Page", fmt.Sprint(0))
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(reviews); err != nil {
		common.SendResponse(w, http.StatusInternalServerError, "Failed to encode response")
	}
}

func (reviewController ReviewController) CreateReview(w http.ResponseWriter, r *http.Request) {
	var requestBody model.ReviewModel
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		common.SendResponse(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	review, err := reviewController.ReviewService.Create(r.Context(), requestBody)
	if err != nil {
		common.SendResponse(w, http.StatusInternalServerError, "Failed to create review")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(review); err != nil {
		common.SendResponse(w, http.StatusInternalServerError, "Failed to encode response")
	}
}
