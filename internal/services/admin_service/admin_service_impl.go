package adminservice

import (
	"context"
	"errors"
	"fmt"
	"strings"

	adminrequest "event_management/internal/dto/request/admin_request"
	"event_management/internal/models"
	adminrepo "event_management/internal/repositories/admin_repositories"
	errorresponse "event_management/pkg/constant/error_response"
	"event_management/pkg/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AdminServiceImpl struct {
	adminRepo adminrepo.IAdminRepository
}

func NewAdminServiceImpl(adminRepo adminrepo.IAdminRepository) IAdminService {
	return &AdminServiceImpl{adminRepo: adminRepo}
}

func (a *AdminServiceImpl) Register(ctx context.Context, req adminrequest.RegisterAdminRequest) error {
	if strings.TrimSpace(req.Name) == "" {
		return errorresponse.NewCustomError(errorresponse.ErrBadRequest, "Name is required", 400)
	}

	if strings.TrimSpace(req.Email) == "" {
		return errorresponse.NewCustomError(errorresponse.ErrBadRequest, "Email is required", 400)
	}

	if strings.TrimSpace(req.Password) == "" {
		return errorresponse.NewCustomError(errorresponse.ErrBadRequest, "Password is required", 400)
	}

	existsEmail, err := a.adminRepo.FindByEmail(ctx, req.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return errorresponse.NewCustomError(errorresponse.ErrInternal, "Failed to get user email", 500)
	}

	if existsEmail != nil {
		return errorresponse.NewCustomError(errorresponse.ErrExists, "Email already exists", 409)
	}

	hashed, err := utils.HashPassword(req.Password)
	if err != nil {
		return errorresponse.NewCustomError(errorresponse.ErrBadRequest, "Failed to hash password", 400)
	}

	role, err := a.adminRepo.FindRoleAdmin(ctx)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errorresponse.NewCustomError(errorresponse.ErrNotFound, "Role 'admin' not found. Please create it first in /api/v1/role/create", 404)
		}
		return errorresponse.NewCustomError(errorresponse.ErrInternal, "Failed to get role admin", 500)
	}

	admin := &models.User{
		ID:       uuid.New(),
		Name:     req.Name,
		Email:    req.Email,
		Password: hashed,
		RoleID:   role.ID,
		Status:   1,
	}

	if err := a.adminRepo.Create(ctx, admin); err != nil {
		return errorresponse.NewCustomError(errorresponse.ErrInternal, "Failed to create admin", 500)
	}

	return nil
}

func (a *AdminServiceImpl) Login(ctx context.Context, req adminrequest.LoginAdminRequest) (string, error) {
	admin, err := a.adminRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		return "", errorresponse.NewCustomError(errorresponse.ErrNotFound, "Invalid credentials", 400)
	}

	isPasswordValid := utils.CheckPasswordHash(req.Password, admin.Password)
	if !isPasswordValid {
		return "", errorresponse.NewCustomError(errorresponse.ErrBadRequest, "Password incorrect", 400)
	}

	token, err := utils.GenerateToken(admin.ID.String(), admin.RoleID.String())
	if err != nil {
		return "", errorresponse.NewCustomError(errorresponse.ErrInternal, "Failed to generate token", 500)
	}

	return token, nil
}

func (a *AdminServiceImpl) GetProfile(ctx context.Context, adminId uuid.UUID) (*models.User, error) {
	admin, err := a.adminRepo.FindByAdminID(ctx, adminId)
	if err != nil {
		return nil, errorresponse.NewCustomError(errorresponse.ErrNotFound, "Admin not found", 404)
	}
	return admin, nil
}

func (a *AdminServiceImpl) UpdateProfile(ctx context.Context, adminID uuid.UUID, req adminrequest.UpdateProfileRequest) error {
	admin, err := a.adminRepo.FindByAdminID(ctx, adminID)
	if err != nil {
		return errorresponse.NewCustomError(errorresponse.ErrNotFound, "Admin not found", 404)
	}

	if strings.TrimSpace(req.Name) == "" {
		return errorresponse.NewCustomError(errorresponse.ErrBadRequest, "Name is required", 400)
	}

	admin.Name = req.Name

	if err := a.adminRepo.Update(ctx, admin); err != nil {
		return errorresponse.NewCustomError(errorresponse.ErrInternal, "Failed to update admin profile", 500)
	}

	return nil
}

func (a *AdminServiceImpl) Logout(ctx context.Context, adminID uuid.UUID) error {
	fmt.Printf("Admin %s logged out\n", adminID.String())
	return nil
}
