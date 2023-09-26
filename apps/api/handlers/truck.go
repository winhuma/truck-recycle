package handlers

import (
	"strconv"
	"truck-unii-app/apps/api/services"
	"truck-unii-app/apps/myconfig/myvar"
	"truck-unii-app/pkg/myresponse"

	"github.com/gofiber/fiber/v2"
)

type hTruck struct {
	st services.ServiceTruck
}

func NewHandlerTruck(st services.ServiceTruck) HandlerTruck {
	return &hTruck{
		st: st,
	}
}

func (h *hTruck) TruckProfileGet(c *fiber.Ctx) error {

	qtruckID := c.Query(myvar.QPTruckID)
	if len(qtruckID) == 0 {
		return c.Status(400).JSON(myresponse.SetResponse(myvar.QPTruckIDMissingMsg))
	}
	truckID, err := strconv.Atoi(qtruckID)
	if err != nil {
		return c.Status(400).JSON(myresponse.SetResponse(myvar.MsgTypeParamWrong))
	}

	data, err := h.st.TruckProfileGet(truckID)
	if err != nil {
		return err
	}

	return c.Status(200).JSON(myresponse.SetResponse(myvar.MsgSuccess, data))
}
