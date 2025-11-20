package standhandler

import (
	standrequest "event_management/internal/dto/request/stand_request"
	standresponse "event_management/internal/dto/response/stand_response"
	standservice "event_management/internal/services/stand_service"
	errorresponse "event_management/pkg/constant/error_response"
	"event_management/pkg/constant/response"
	"event_management/pkg/utils"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type StandHandler struct {
	standService standservice.IStandService
}

func NewStandHandler(standService standservice.IStandService) *StandHandler {
	return &StandHandler{standService: standService}
}

func (s *StandHandler) CreateStand(c echo.Context) error {
	var req standrequest.CreateStandRequest
	req.StandName = c.FormValue("stand_name")
	req.Address = c.FormValue("address")

	lat, _ := strconv.ParseFloat(c.FormValue("lat"), 64)
	lng, _ := strconv.ParseFloat(c.FormValue("lng"), 64)
	req.Lat = lat
	req.Lng = lng

	standCategoryID := c.FormValue("stand_category_id")
	if standCategoryID != "" {
		id, err := uuid.Parse(standCategoryID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "Invalid stand_category_id")
		}
		req.StandCategoryID = id
	}

	err := s.standService.CreateStand(c.Request().Context(), req)
	if err != nil {
		if customErr, ok := errorresponse.AsCustomErr(err); ok {
			return response.Error(c, customErr.Status, customErr.Msg, customErr.Err.Error())
		}
		return response.Error(c, http.StatusInternalServerError, err.Error(), "failed to create role")
	}

	return c.JSON(http.StatusCreated, "Stand Created Successfully")
}

func (s *StandHandler) GetAllStand(c echo.Context) error {
	pageInt, limitInt := utils.ParsePaginationParams(c, 10)
	search := c.QueryParam("search")

	stand, total, err := s.standService.GetAllStand(c.Request().Context(), pageInt, limitInt, search)
	if err != nil {
		if customErr, ok := errorresponse.AsCustomErr(err); ok {
			return response.Error(c, customErr.Status, customErr.Msg, customErr.Err.Error())
		}
		return response.Error(c, http.StatusInternalServerError, err.Error(), "failed to get stand category")
	}

	meta := utils.BuildPaginationMeta(c, pageInt, limitInt, total)
	data := make([]standresponse.StandResponse, len(stand))
	for i, s := range stand {
		data[i] = standresponse.ToStandResponse(*s)
	}

	return response.PaginatedSuccess(c, http.StatusOK, "Get All Stand Successfully", data, meta)
}

func (s *StandHandler) GetByIdStand(c echo.Context) error {
	standId, err := uuid.Parse(c.Param("standId"))
	if err != nil {
		return response.Error(c, http.StatusBadRequest, "bad request", err.Error())
	}

	stand, err := s.standService.GetByIdStand(c.Request().Context(), standId)
	if err != nil {
		if customErr, ok := errorresponse.AsCustomErr(err); ok {
			return response.Error(c, customErr.Status, customErr.Msg, customErr.Err.Error())
		}
		return response.Error(c, http.StatusInternalServerError, err.Error(), "failed to get stand")
	}

	res := standresponse.ToStandResponse(*stand)

	return response.Success(c, http.StatusOK, "Get Stand Successfully", res)
}

func (s *StandHandler) UpdateStand(c echo.Context) error {
	standId, err := uuid.Parse(c.Param("standId"))
	if err != nil {
		return response.Error(c, http.StatusBadRequest, "bad request", err.Error())
	}

	var req standrequest.UpdateStandRequest
	if err := c.Bind(&req); err != nil {
		return response.Error(c, http.StatusBadRequest, "bad request", err.Error())
	}

	err = s.standService.UpdateStand(c.Request().Context(), standId, req)
	if err != nil {
		if customErr, ok := errorresponse.AsCustomErr(err); ok {
			return response.Error(c, customErr.Status, customErr.Msg, customErr.Err.Error())
		}
		return response.Error(c, http.StatusInternalServerError, err.Error(), "failed to update stand")
	}

	return response.Success(c, http.StatusOK, "Stand Updated Successfully", nil)
}

func (s *StandHandler) DeleteStand(c echo.Context) error {
	standId, err := uuid.Parse(c.Param("standId"))
	if err != nil {
		return response.Error(c, http.StatusBadRequest, "bad request", err.Error())
	}

	if err := s.standService.DeleteStand(c.Request().Context(), standId); err != nil {
		if customErr, ok := errorresponse.AsCustomErr(err); ok {
			return response.Error(c, customErr.Status, customErr.Msg, customErr.Err.Error())
		}
		return response.Error(c, http.StatusInternalServerError, err.Error(), "failed to delete stand")
	}

	return response.Success(c, http.StatusOK, "Stand Deleted Successfully", nil)
}
