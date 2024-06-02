package eventcontroller

import (
	"net/http"
	"strconv"

	"github.com/ArdiSasongko/app_ticketing/db/model/web"
	"github.com/ArdiSasongko/app_ticketing/helper"
	eventservice "github.com/ArdiSasongko/app_ticketing/service/event.service"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type EventController struct {
	service eventservice.EventServiceInterface
}

func NewEventController(service eventservice.EventServiceInterface) *EventController {
	return &EventController{
		service: service,
	}
}

func (con *EventController) Create(c echo.Context) error {
	event := new(web.EventRequest)
	user := c.Get("user").(*jwt.Token)
	claims, _ := user.Claims.(*helper.CustomClaims)
	if err := c.Bind(event); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}

	if err := c.Validate(event); err != nil {
		return err
	}

	result, err := con.service.Create(claims.UserID, *event)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}

	return c.JSON(http.StatusCreated, helper.ResponseToClient(http.StatusCreated, "Success create event", result))
}

func (con *EventController) FetchAll(c echo.Context) error {
	result, err := con.service.FetchAll()

	if err != nil {
		return c.JSON(http.StatusNotFound, helper.ResponseToClient(http.StatusNotFound, err.Error(), nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseToClient(http.StatusOK, "Success fetch all event", result))
}

func (con *EventController) FetchEvent(c echo.Context) error {
	eventID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}

	result, errFetch := con.service.FetchEvent(eventID)

	if errFetch != nil {
		return c.JSON(http.StatusNotFound, helper.ResponseToClient(http.StatusNotFound, errFetch.Error(), nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseToClient(http.StatusOK, "Success fetch event", result))
}

func (con *EventController) UpdateEvent(c echo.Context) error {
	eventID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}

	eventUpdate := new(web.EventUpdateRequest)

	if err := c.Bind(eventUpdate); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}

	if err := c.Validate(eventUpdate); err != nil {
		return err
	}

	result, errUpdate := con.service.UpdateEvent(eventID, *eventUpdate)

	if errUpdate != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseToClient(http.StatusBadRequest, errUpdate.Error(), nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseToClient(http.StatusOK, "Success update event", result))
}
