package instancehandler

import (
	instancerequest "event_management/internal/dto/request/instance_request"
	instanceresponse "event_management/internal/dto/response/instance_response"
	instanceservice "event_management/internal/services/instance_service"
	errorresponse "event_management/pkg/constant/error_response"
	"event_management/pkg/constant/response"
	"event_management/pkg/utils"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type InstanceHandler struct {
	instanceService instanceservice.IInstanceService
}

func NewInstanceHandler(instanceService instanceservice.IInstanceService) *InstanceHandler {
	return &InstanceHandler{instanceService: instanceService}
}

func (r *InstanceHandler) CreateInstance(c echo.Context) error {
	var req instancerequest.CreateInstanceRequest
	if err := c.Bind(&req); err != nil {
		return response.Error(c, http.StatusBadRequest, "bad request", err.Error())
	}

	err := r.instanceService.CreateInstance(c.Request().Context(), req)
	if err != nil {
		if customErr, ok := errorresponse.AsCustomErr(err); ok {
			return response.Error(c, customErr.Status, customErr.Msg, customErr.Err.Error())
		}
		return response.Error(c, http.StatusInternalServerError, err.Error(), "failed to create instance")
	}

	return response.Success(c, http.StatusOK, "Instance Created Successfully", nil)
}

func (r *InstanceHandler) GetAllInstance(c echo.Context) error {
	pageInt, limitInt := utils.ParsePaginationParams(c, 10)
	search := c.QueryParam("search")

	instances, total, err := r.instanceService.GetAllInstance(c.Request().Context(), pageInt, limitInt, search)
	if err != nil {
		if customErr, ok := errorresponse.AsCustomErr(err); ok {
			return response.Error(c, customErr.Status, customErr.Msg, customErr.Err.Error())
		}
		return response.Error(c, http.StatusInternalServerError, err.Error(), "failed to get instance")
	}

	meta := utils.BuildPaginationMeta(c, pageInt, limitInt, total)
	data := make([]instanceresponse.InstanceResponse, len(instances))
	for i, s := range instances {
		data[i] = instanceresponse.ToInstanceResponse(*s)
	}

	return response.PaginatedSuccess(c, http.StatusOK, "Get All Instance Successfully", data, meta)
}

func (r *InstanceHandler) GetByIdInstance(c echo.Context) error {
	instanceId, err := uuid.Parse(c.Param("instanceId"))
	if err != nil {
		return response.Error(c, http.StatusBadRequest, "bad request", err.Error())
	}

	instance, err := r.instanceService.GetByIdInstance(c.Request().Context(), instanceId)
	if err != nil {
		if customErr, ok := errorresponse.AsCustomErr(err); ok {
			return response.Error(c, customErr.Status, customErr.Msg, customErr.Err.Error())
		}
		return response.Error(c, http.StatusInternalServerError, err.Error(), "failed to get instance")
	}

	res := instanceresponse.ToInstanceResponse(*instance)

	return response.Success(c, http.StatusOK, "Get Instance Successfully", res)
}

func (r *InstanceHandler) UpdateInstance(c echo.Context) error {
	instanceId, err := uuid.Parse(c.Param("instanceId"))
	if err != nil {
		return response.Error(c, http.StatusBadRequest, "bad request", err.Error())
	}

	var req instancerequest.UpdateInstanceRequest
	if err := c.Bind(&req); err != nil {
		return response.Error(c, http.StatusBadRequest, "bad request", err.Error())
	}

	err = r.instanceService.UpdateInstance(c.Request().Context(), instanceId, req)
	if err != nil {
		if customErr, ok := errorresponse.AsCustomErr(err); ok {
			return response.Error(c, customErr.Status, customErr.Msg, customErr.Err.Error())
		}
		return response.Error(c, http.StatusInternalServerError, err.Error(), "failed to update instance")
	}

	return response.Success(c, http.StatusOK, "Instance Updated Successfully", nil)
}

func (r *InstanceHandler) DeleteInstance(c echo.Context) error {
	instanceId, err := uuid.Parse(c.Param("instanceId"))
	if err != nil {
		return response.Error(c, http.StatusBadRequest, "bad request", err.Error())
	}

	if err := r.instanceService.DeleteInstance(c.Request().Context(), instanceId); err != nil {
		if customErr, ok := errorresponse.AsCustomErr(err); ok {
			return response.Error(c, customErr.Status, customErr.Msg, customErr.Err.Error())
		}
		return response.Error(c, http.StatusInternalServerError, err.Error(), "failed to delete instance")
	}

	return response.Success(c, http.StatusOK, "Instance Deleted Successfully", nil)
}