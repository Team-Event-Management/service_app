package standcategoryresponse

import (
	"event_management/internal/models"
	"event_management/pkg/utils"

	"github.com/google/uuid"
)

type StandCategoryResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   string    `json:"created_at"`
	UpdatedAt   string    `json:"updated_at"`
}

func ToStandCategoryResponse(stand_category models.StandCategory) StandCategoryResponse {
	return StandCategoryResponse{
		ID:          stand_category.ID,
		Name:        stand_category.Name,
		Description: stand_category.Description,
		CreatedAt:   utils.FormatDate(stand_category.CreatedAt),
		UpdatedAt:   utils.FormatDate(stand_category.UpdatedAt),
	}
}
