package standresponse

import (
	"event_management/internal/models"
	"event_management/pkg/utils"

	"github.com/google/uuid"
)

type StandResponse struct {
	ID              uuid.UUID `json:"id"`
	StandName       string    `json:"stand_name"`
	Lat             float64   `json:"lat"`
	Lng             float64   `json:"lng"`
	Address         string    `json:"address"`
	StandCategoryID uuid.UUID `json:"stand_category_id"`
	CreatedAt       string    `json:"created_at"`
	UpdatedAt       string    `json:"updated_at"`
}

func ToStandResponse(stand models.Stand) StandResponse {
	return StandResponse{
		ID:              stand.ID,
		StandName:       stand.StandName,
		Lat:             stand.Lat,
		Lng:             stand.Lng,
		Address:         stand.Address,
		StandCategoryID: stand.StandCategoryID,
		CreatedAt:       utils.FormatDate(stand.CreatedAt),
		UpdatedAt:       utils.FormatDate(stand.UpdatedAt),
	}
}
