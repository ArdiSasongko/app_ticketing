package ordercontroller

import (
	"net/http"
	"strconv"

	"github.com/ArdiSasongko/app_ticketing/db/model/web"
	"github.com/ArdiSasongko/app_ticketing/helper"
	orderservice "github.com/ArdiSasongko/app_ticketing/service/order.service"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type OrderController struct {
	service orderservice.OrderServiceInterface
}

func NewOrderController(service orderservice.OrderServiceInterface) *OrderController {
	return &OrderController{
		service: service,
	}
}

func (con *OrderController) Create(c echo.Context) error {
	eventID, err := strconv.Atoi(c.Param("event_id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}

	ticketID, err := strconv.Atoi(c.Param("ticket_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}

	user := c.Get("user").(*jwt.Token)
	claims, _ := user.Claims.(*helper.CustomClaims)
	userID := claims.UserID

	orderReq := new(web.OrderRequest)
	if err := c.Bind(orderReq); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}

	if err := c.Validate(orderReq); err != nil {
		return err
	}

	orderReq.EventID = eventID
	orderReq.TicketID = ticketID
	orderReq.BuyerID = userID

	result, err := con.service.Create(*orderReq)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseToClient(http.StatusBadRequest, err.Error(), nil))
	}

	return c.JSON(http.StatusOK, helper.ResponseToClient(http.StatusOK, "Success", result))
}
