package eventresponse

import (
	"event_management/internal/models"
	"event_management/pkg/utils"

	"github.com/google/uuid"
)

type ImageResponse struct {
	ID        uuid.UUID `json:"id"`
	ImagePath string    `json:"image_path"`
}

type EventResponse struct {
	ID           uuid.UUID       `json:"id"`
	NameEvent    string          `json:"name_event"`
	Description  string          `json:"description"`
	Status       int             `json:"status"`
	Location     string          `json:"location"`
	StartDate 	 string 		 `json:"start_date"`
	EventImages  []ImageResponse `json:"event_images"`
	CreatedAt    string          `json:"created_at"`
	UpdatedAt    string          `json:"updated_at"`
}

func ToEventResponse(event models.Event) EventResponse {
	images := make([]ImageResponse, 0)
	for _, img := range event.EventImages {
		images = append(images, ImageResponse{
			ID:        img.ID,
			ImagePath: img.ImagePath,
		})
	}

	return EventResponse{
		ID:           event.ID,
		NameEvent:    event.NameEvent,
		Description:  event.Description,
		Status:       event.Status,
		Location:     event.Location,
		StartDate: 	  utils.FormatDate(event.StartDate),
		EventImages:  images,
		CreatedAt:    utils.FormatDate(event.CreatedAt),
		UpdatedAt:    utils.FormatDate(event.UpdatedAt),
	}
}
