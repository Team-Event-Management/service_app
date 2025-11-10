package instanceresponse

import (
	"event_management/internal/models"
	"event_management/pkg/utils"

	"github.com/google/uuid"
)

type InstanceResponse struct {
	ID        	uuid.UUID `json:"id"`
	Name	  	string    `json:"name"`
	Lat 	 	float64   `json:"lat"`
	Lng 	 	float64   `json:"lng"`
	FullAddress string    `json:"full_address"`
	CreatedAt 	string    `json:"created_at"`
	UpdatedAt 	string    `json:"updated_at"`
}

func ToInstanceResponse(instance models.Instance) InstanceResponse {
	return InstanceResponse{
		ID:        	 instance.ID,
		Name:      	 instance.Name,
		Lat:       	 instance.Lat,
		Lng:       	 instance.Lng,
		FullAddress: instance.FullAddress,
		CreatedAt:   utils.FormatDate(instance.CreatedAt),
		UpdatedAt:   utils.FormatDate(instance.UpdatedAt),
	}
}