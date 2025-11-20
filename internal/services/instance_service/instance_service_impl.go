package instanceservice

import (
	"context"
	instancerequest "event_management/internal/dto/request/instance_request"
	"event_management/internal/models"
	instancerepository "event_management/internal/repositories/instance_repositories"
	errorresponse "event_management/pkg/constant/error_response"
	"strings"

	"github.com/google/uuid"
)

type InstanceServiceImpl struct {
	instanceRepo instancerepository.IInstanceRepository
}

func NewInstanceServiceImpl(instanceRepo instancerepository.IInstanceRepository) IInstanceService {
	return &InstanceServiceImpl{instanceRepo: instanceRepo}
}

func (s *InstanceServiceImpl) CreateInstance(ctx context.Context, req instancerequest.CreateInstanceRequest) error {
	if strings.TrimSpace(req.Name) == "" {
		return errorresponse.NewCustomError(errorresponse.ErrBadRequest, "Instance Name is required", 400)
	}

	instance := &models.Instance{
		Name:        req.Name,
		Lat:         req.Lat,
		Lng:         req.Lng,
		FullAddress: req.FullAddress,
	}
	
	return s.instanceRepo.Create(ctx, instance)
}

func (s *InstanceServiceImpl) GetAllInstance(ctx context.Context, page, limit int, search string) ([]*models.Instance, int, error) {
	offset := (page - 1) * limit
	return s.instanceRepo.FindAll(ctx, limit, offset, search)
}

func (s *InstanceServiceImpl) GetByIdInstance(ctx context.Context, instanceId uuid.UUID) (*models.Instance, error) {
	return s.instanceRepo.FindById(ctx, instanceId)
}

func (s *InstanceServiceImpl) UpdateInstance(ctx context.Context, instanceId uuid.UUID, req instancerequest.UpdateInstanceRequest) error {
	instance, err := s.instanceRepo.FindById(ctx, instanceId)
	if err != nil {
		return err
	}

	if req.Name != "" {
		instance.Name = req.Name
	}
	instance.Lat = req.Lat
	instance.Lng = req.Lng
	if req.FullAddress != "" {
		instance.FullAddress = req.FullAddress
	}

	return s.instanceRepo.Update(ctx, instanceId, instance)
}

func (s *InstanceServiceImpl) DeleteInstance(ctx context.Context, instanceId uuid.UUID) error {
	return s.instanceRepo.Delete(ctx, instanceId)
}
