package controller

import (
	"encoding/json"
	"guitarStore/api"
	"guitarStore/internal/common"
	"guitarStore/internal/service"
	"net/http"
)

type PromocodeController struct {
	service.PromocodeService
}

func NewPromocodeController(promocodeService *service.PromocodeService) *PromocodeController {
	return &PromocodeController{PromocodeService: *promocodeService}
}

func (promocodeController PromocodeController) GetPromocodeByCode(w http.ResponseWriter, r *http.Request, params api.GetPromocodeByCodeParams) {
	ctx := r.Context()

	promocode, err := promocodeController.PromocodeService.FindByCode(ctx, params.Code)
	if err != nil {
		common.SendResponse(w, http.StatusInternalServerError, "Failed to retrieve promocode")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(promocode); err != nil {
		common.SendResponse(w, http.StatusInternalServerError, "Failed to encode response")
	}
}
