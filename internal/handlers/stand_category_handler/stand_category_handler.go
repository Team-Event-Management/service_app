package standcategoryhandler

import (
	standcategoryrequest "event_management/internal/dto/request/stand_category_request"
	standcategoryresponse "event_management/internal/dto/response/stand_category_response"
	standcategoryservice "event_management/internal/services/stand_category_service"
	errorresponse "event_management/pkg/constant/error_response"
	"event_management/pkg/constant/response"
	"event_management/pkg/utils"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type StandCategoryHandler struct {
	standCategoryService standcategoryservice.IStandCategoryService
}

func NewStandCategoryHandler(standCategoryService standcategoryservice.IStandCategoryService) *StandCategoryHandler {
	return &StandCategoryHandler{standCategoryService: standCategoryService}
}

func (s *StandCategoryHandler) CreateStandCategory(c echo.Context) error {
	var req standcategoryrequest.CreateStandCategoryRequest
	req.Name = c.FormValue("name")
	req.Description = c.FormValue("description")

	if req.Name == "" {
		return response.Error(c, 400, "name is required", "empty name")
	}

	err := s.standCategoryService.CreateStandCategory(c.Request().Context(), req)
	if err != nil {
		if customErr, ok := errorresponse.AsCustomErr(err); ok {
			return response.Error(c, customErr.Status, customErr.Msg, customErr.Err.Error())
		}
		return response.Error(c, http.StatusInternalServerError, err.Error(), "failed to create role")
	}

	return c.JSON(http.StatusCreated, "Stand Category Created Successfully")
}

func (s *StandCategoryHandler) GetAllStandCategory(c echo.Context) error {
	pageInt, limitInt := utils.ParsePaginationParams(c, 10)
	search := c.QueryParam("search")

	standCategories, total, err := s.standCategoryService.GetAllStandCategory(c.Request().Context(), pageInt, limitInt, search)
	if err != nil {
		if customErr, ok := errorresponse.AsCustomErr(err); ok {
			return response.Error(c, customErr.Status, customErr.Msg, customErr.Err.Error())
		}
		return response.Error(c, http.StatusInternalServerError, err.Error(), "failed to get stand category")
	}

	meta := utils.BuildPaginationMeta(c, pageInt, limitInt, total)
	data := make([]standcategoryresponse.StandCategoryResponse, len(standCategories))
	for i, s := range standCategories {
		data[i] = standcategoryresponse.ToStandCategoryResponse(*s)
	}

	return response.PaginatedSuccess(c, http.StatusOK, "Get All Stand Category Successfully", data, meta)
}

func (s *StandCategoryHandler) GetByIdStandCategory(c echo.Context) error {
	standCategoryId, err := uuid.Parse(c.Param("standCategoryId"))
	if err != nil {
		return response.Error(c, http.StatusBadRequest, "bad request", err.Error())
	}

	standCategory, err := s.standCategoryService.GetByIdStandCategory(c.Request().Context(), standCategoryId)
	if err != nil {
		if customErr, ok := errorresponse.AsCustomErr(err); ok {
			return response.Error(c, customErr.Status, customErr.Msg, customErr.Err.Error())
		}
		return response.Error(c, http.StatusInternalServerError, err.Error(), "failed to get stand category")
	}

	res := standcategoryresponse.ToStandCategoryResponse(*standCategory)

	return response.Success(c, http.StatusOK, "Get Stand Category Successfully", res)
}

func (s *StandCategoryHandler) UpdateStandCategory(c echo.Context) error {
	standCategoryId, err := uuid.Parse(c.Param("standCategoryId"))
	if err != nil {
		return response.Error(c, http.StatusBadRequest, "bad request", err.Error())
	}

	var req standcategoryrequest.UpdateStandCategoryRequest
	if err := c.Bind(&req); err != nil {
		return response.Error(c, http.StatusBadRequest, "bad request", err.Error())
	}

	err = s.standCategoryService.UpdateStandCategory(c.Request().Context(), standCategoryId, req)
	if err != nil {
		if customErr, ok := errorresponse.AsCustomErr(err); ok {
			return response.Error(c, customErr.Status, customErr.Msg, customErr.Err.Error())
		}
		return response.Error(c, http.StatusInternalServerError, "failed to update event", err.Error())
	}

	return response.Success(c, http.StatusOK, "Stand Category Updated Successfully", nil)
}

func (s *StandCategoryHandler) DeleteStandCategory(c echo.Context) error {
	standCategoryId, err := uuid.Parse(c.Param("standCategoryId"))
	if err != nil {
		return response.Error(c, http.StatusBadRequest, "bad request", err.Error())
	}

	if err := s.standCategoryService.DeleteStandCategory(c.Request().Context(), standCategoryId); err != nil {
		if customErr, ok := errorresponse.AsCustomErr(err); ok {
			return response.Error(c, customErr.Status, customErr.Msg, customErr.Err.Error())
		}
		return response.Error(c, http.StatusInternalServerError, err.Error(), "failed to delete stand category")
	}

	return response.Success(c, http.StatusOK, "Stand Category Deleted Successfully", nil)
}
