package eventhandler

import (
	eventrequest "event_management/internal/dto/request/event_request"
	eventresponse "event_management/internal/dto/response/event_response"
	eventservice "event_management/internal/services/event_service"
	imageservice "event_management/internal/services/image_service"
	errorresponse "event_management/pkg/constant/error_response"
	"event_management/pkg/constant/response"
	"event_management/pkg/utils"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type EventHandler struct {
	eventService eventservice.IEventService
	imageService imageservice.IImageService
}

func NewEventHandler(
	eventService eventservice.IEventService,
	imageService imageservice.IImageService,
) *EventHandler {
	return &EventHandler{
		eventService: eventService,
		imageService: imageService,
	}
}

func (r *EventHandler) CreateEvent(c echo.Context) error {
	var req eventrequest.CreateEventRequest
	if err := c.Bind(&req); err != nil {
		return response.Error(c, http.StatusBadRequest, "bad request", err.Error())
	}

	err := r.eventService.CreateEvent(c.Request().Context(), req)
	if err != nil {
		if customErr, ok := errorresponse.AsCustomErr(err); ok {
			return response.Error(c, customErr.Status, customErr.Msg, customErr.Err.Error())
		}
		return response.Error(c, http.StatusInternalServerError, "failed to create event", err.Error())
	}

	return response.Success(c, http.StatusCreated, "Event Created Successfully", nil)
}

func (r *EventHandler) GetAllEvent(c echo.Context) error {
	page, limit := utils.ParsePaginationParams(c, 10)
	search := c.QueryParam("search")

	events, total, err := r.eventService.GetAllEvent(c.Request().Context(), page, limit, search)
	if err != nil {
		if customErr, ok := errorresponse.AsCustomErr(err); ok {
			return response.Error(c, customErr.Status, customErr.Msg, customErr.Err.Error())
		}
		return response.Error(c, http.StatusInternalServerError, "failed to get events", err.Error())
	}

	meta := utils.BuildPaginationMeta(c, page, limit, total)
	data := make([]eventresponse.EventResponse, len(events))
	for i, s := range events {
		data[i] = eventresponse.ToEventResponse(*s)
	}

	return response.PaginatedSuccess(c, http.StatusOK, "Get All Event Successfully", data, meta)
}

func (r *EventHandler) GetByIdEvent(c echo.Context) error {
	eventId, err := uuid.Parse(c.Param("eventId"))
	if err != nil {
		return response.Error(c, http.StatusBadRequest, "bad request", err.Error())
	}

	event, err := r.eventService.GetByIdEvent(c.Request().Context(), eventId)
	if err != nil {
		if customErr, ok := errorresponse.AsCustomErr(err); ok {
			return response.Error(c, customErr.Status, customErr.Msg, customErr.Err.Error())
		}
		return response.Error(c, http.StatusInternalServerError, err.Error(), "failed to get event	")
	}

	res := eventresponse.ToEventResponse(*event)

	return response.Success(c, http.StatusOK, "Get Event Successfully", res)
}
	
func (r *EventHandler) UpdateEvent(c echo.Context) error {
	eventId, err := uuid.Parse(c.Param("eventId"))
	if err != nil {
		return response.Error(c, http.StatusBadRequest, "bad request", err.Error())
	}

	var req eventrequest.UpdateEventRequest
	if err := c.Bind(&req); err != nil {
		return response.Error(c, http.StatusBadRequest, "bad request", err.Error())
	}

	err = r.eventService.UpdateEvent(c.Request().Context(), eventId, req)
	if err != nil {
		if customErr, ok := errorresponse.AsCustomErr(err); ok {
			return response.Error(c, customErr.Status, customErr.Msg, customErr.Err.Error())
		}
		return response.Error(c, http.StatusInternalServerError, "failed to update event", err.Error())
	}

	return response.Success(c, http.StatusOK, "Event Updated Successfully", nil)
}

func (r *EventHandler) DeleteEvent(c echo.Context) error {
	eventId, err := uuid.Parse(c.Param("eventId"))
	if err != nil {
		return response.Error(c, http.StatusBadRequest, "bad request", err.Error())
	}

	err = r.eventService.DeleteEvent(c.Request().Context(), eventId)
	if err != nil {
		if customErr, ok := errorresponse.AsCustomErr(err); ok {
			return response.Error(c, customErr.Status, customErr.Msg, customErr.Err.Error())
		}
		return response.Error(c, http.StatusInternalServerError, "failed to delete event", err.Error())
	}

	return response.Success(c, http.StatusOK, "Event Deleted Successfully", nil)
}

func (r *EventHandler) UploadEventImage(c echo.Context) error {
	eventId, err := uuid.Parse(c.Param("eventId"))
	if err != nil {
		return response.Error(c, http.StatusBadRequest, "bad request", err.Error())
	}

	file, err := c.FormFile("file")
	if err != nil {
		return response.Error(c, http.StatusBadRequest, "bad request", err.Error())
	}

	_, err = r.imageService.UploadImage(c.Request().Context(), eventId, file)
	if err != nil {
		if customErr, ok := errorresponse.AsCustomErr(err); ok {
			return response.Error(c, customErr.Status, customErr.Msg, customErr.Err.Error())
		}
		return response.Error(c, http.StatusInternalServerError, err.Error(), "failed to upload image")
	}

	return response.Success(c, http.StatusCreated, "Image Uploaded Successfully", nil)
}

func (r *EventHandler) ListEventImages(c echo.Context) error {
	eventId, err := uuid.Parse(c.Param("eventId"))
	if err != nil {
		return response.Error(c, http.StatusBadRequest, "bad request", err.Error())
	}

	images, err := r.imageService.ListImages(c.Request().Context(), eventId)
	if err != nil {
		if customErr, ok := errorresponse.AsCustomErr(err); ok {
			return response.Error(c, customErr.Status, customErr.Msg, customErr.Err.Error())
		}
		return response.Error(c, http.StatusInternalServerError, err.Error(), "failed to get images")
	}

	return response.Success(c, http.StatusOK, "Get All Images Successfully", images)
}

func (r *EventHandler) DeleteEventImage(c echo.Context) error {
	eventId, err := uuid.Parse(c.Param("eventId"))
	if err != nil {
		return response.Error(c, http.StatusBadRequest, "bad request", err.Error())
	}

	imageID, err := uuid.Parse(c.Param("imageId"))
	if err != nil {
		return response.Error(c, http.StatusBadRequest, "bad request", err.Error())
	}

	err = r.imageService.DeleteImage(c.Request().Context(), eventId, imageID)
	if err != nil {
		if customErr, ok := errorresponse.AsCustomErr(err); ok {
			return response.Error(c, customErr.Status, customErr.Msg, customErr.Err.Error())
		}
		return response.Error(c, http.StatusInternalServerError, err.Error(), "failed to delete image")
	}

	return response.Success(c, http.StatusOK, "Image Deleted Successfully", nil)
}