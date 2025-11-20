package standservice

import (
	"context"
	standrequest "event_management/internal/dto/request/stand_request"
	"event_management/internal/models"
	standrepositories "event_management/internal/repositories/stand_repositories"
	errorresponse "event_management/pkg/constant/error_response"
	"strings"

	"github.com/google/uuid"
)

type StandServiceImpl struct {
	standRepo standrepositories.IStandRepository
}

func NewStandServiceImpl(standRepo standrepositories.IStandRepository) IStandService {
	return &StandServiceImpl{standRepo: standRepo}
}

func (s *StandServiceImpl) CreateStand(ctx context.Context, req standrequest.CreateStandRequest) error {
	if strings.TrimSpace(req.StandName) == "" {
		return errorresponse.NewCustomError(errorresponse.ErrBadRequest, "Stand Name is required", 400)
	}

	stand := &models.Stand{
		StandName:       req.StandName,
		Lat:             req.Lat,
		Lng:             req.Lng,
		Address:         req.Address,
		StandCategoryID: req.StandCategoryID,
	}

	return s.standRepo.Create(ctx, stand)
}

func (s *StandServiceImpl) GetAllStand(ctx context.Context, page, limit int, search string) ([]*models.Stand, int, error) {
	offset := (page - 1) * limit
	return s.standRepo.FindAll(ctx, limit, offset, search)
}

func (s *StandServiceImpl) GetByIdStand(ctx context.Context, standId uuid.UUID) (*models.Stand, error) {
	return s.standRepo.FindById(ctx, standId)
}

func (s *StandServiceImpl) UpdateStand(ctx context.Context, standId uuid.UUID, req standrequest.UpdateStandRequest) error {
	stand, err := s.standRepo.FindById(ctx, standId)
	if err != nil {
		return err
	}

	if req.StandName != "" {
		stand.StandName = req.StandName
	}

	if req.Lat != 0 {
		stand.Lat = req.Lat
	}

	if req.Lng != 0 {
		stand.Lng = req.Lng
	}

	if req.Address != "" {
		stand.Address = req.Address
	}

	if req.StandCategoryID != uuid.Nil {
		stand.StandCategoryID = req.StandCategoryID
	}

	return s.standRepo.Update(ctx, standId, stand)
}

func (s *StandServiceImpl) DeleteStand(ctx context.Context, standId uuid.UUID) error {
	return s.standRepo.Delete(ctx, standId)
}
