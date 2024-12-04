package controller

import (
	"encoding/json"
	"fmt"
	openapiTypes "github.com/oapi-codegen/runtime/types"
	"guitarStore/api"
	"guitarStore/internal/common"
	"guitarStore/internal/service"
	"net/http"
)

type GuitarController struct {
	service.GuitarService
}

func NewGuitarController(guitarService *service.GuitarService) *GuitarController {
	return &GuitarController{GuitarService: *guitarService}
}

func (guitarController GuitarController) FindAllGuitars(w http.ResponseWriter, r *http.Request, params api.FindAllGuitarsParams) {
	ctx := r.Context()

	guitars, err := guitarController.GuitarService.FindAll(ctx)
	if err != nil {
		common.SendResponse(w, http.StatusInternalServerError, "Failed to retrieve guitars")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Pagination-Count", fmt.Sprint(0))
	w.Header().Set("X-Pagination-Limit", fmt.Sprint(0))
	w.Header().Set("X-Pagination-Page", fmt.Sprint(0))
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(guitars); err != nil {
		common.SendResponse(w, http.StatusInternalServerError, "Failed to encode response")
	}
}

func (guitarController GuitarController) GetGuitarById(w http.ResponseWriter, r *http.Request, guitarId openapiTypes.UUID) {
	ctx := r.Context()

	guitar, err := guitarController.GuitarService.FindById(ctx, guitarId)
	if err != nil {
		common.SendResponse(w, http.StatusInternalServerError, "Failed to retrieve guitar")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(guitar); err != nil {
		common.SendResponse(w, http.StatusInternalServerError, "Failed to encode response")
	}
}
