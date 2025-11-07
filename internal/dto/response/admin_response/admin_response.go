package adminresponse

import (
	"event_management/internal/models"
	"event_management/pkg/utils"

	"github.com/google/uuid"
)

type AdminResponse struct {
	ID        uuid.UUID `json:"id"`
	Name	  string    `json:"name"`
	Email	  string    `json:"email"`
	Status    int       `json:"status"`
	Role      string    `json:"role"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
}

func ToAdminResponse(admin models.User) AdminResponse {
	return AdminResponse{
		ID:        admin.ID,
		Name:      admin.Name,
		Email:     admin.Email,
		Status:    admin.Status,
		Role:      admin.Role.Name,
		CreatedAt: utils.FormatDate(admin.CreatedAt),
		UpdatedAt: utils.FormatDate(admin.UpdatedAt),
	}
}