package handlerconsumer

import (
	"context"
	"event_management/configs"
	"event_management/internal/models"
	eventrepo "event_management/internal/repositories/event_repositories"
	imagerepo "event_management/internal/repositories/image_repositories"
	"event_management/pkg/workers/payload"
	"log"
)

// EventImageHandler menangani penyimpanan image hasil upload ke Cloudinary
type EventImageHandler struct{}

func (h *EventImageHandler) HandleSingle(ctx context.Context, imageURL string, payloads any) error {
	log.Println("⚠️ HandleSingle called in EventImageHandler — not implemented (ignored).")
	return nil
}

// HandleMany menyimpan data image dan mengaitkannya ke event
func (h *EventImageHandler) HandleMany(ctx context.Context, image *models.Image, payloads any) error {
	p, ok := payloads.(*payload.ImageUploadPayload)
	if !ok {
		log.Println("⚠️ payload type mismatch in EventImageHandler")
		return nil
	}

	// Inisialisasi repository
	eventRepo := eventrepo.NewEventRepositoryImpl(configs.DB)
	imageRepo := imagerepo.NewImageRepositoryImpl(configs.DB)

	// Pastikan event-nya ada
	event, err := eventRepo.FindById(ctx, p.ID)
	if err != nil {
		log.Printf("❌ Event with ID %s not found: %v", p.ID, err)
		return err
	}

	// Simpan image ke tabel images
	if err := imageRepo.Create(ctx, image); err != nil {
		log.Printf("❌ Failed to create image record: %v", err)
		return err
	}

	// Kaitkan image ke event (tabel pivot image_events)
	if err := imageRepo.AttachToEvent(ctx, event.ID, image.ID); err != nil {
		log.Printf("❌ Failed to attach image to event: %v", err)
		return err
	}

	log.Printf("✅ Image %s attached successfully to event %s", image.ImagePath, event.ID)
	return nil
}
