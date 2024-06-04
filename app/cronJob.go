package app

import (
	"log"

	orderservice "github.com/ArdiSasongko/app_ticketing/service/order.service"
	"github.com/robfig/cron/v3"
)

func InitCron(service orderservice.OrderServiceInterface) *cron.Cron {
	c := cron.New()

	// running every 5 minutes
	c.AddFunc("@every 5m", func() {
		err := service.CanceledOrder()
		if err != nil {
			log.Printf("Error cancelling expired orders: %v", err)
		} else {
			log.Println("Successfully cancelled expired orders")
		}
	})

	c.Start()
	return c
}
