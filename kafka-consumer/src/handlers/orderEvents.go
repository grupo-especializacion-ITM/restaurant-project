package handlers

import (
	"encoding/json"
	"fmt"
	"log"

	kitlog "github.com/go-kit/log"
	"kafka.consumer.go/src/config"
	"kafka.consumer.go/src/email"
	"kafka.consumer.go/src/models"
)

type OrderHandler struct {
	Cfg *config.Conf
	Log *kitlog.Logger
}

func NewOrderHandler(cfg *config.Conf, logger *kitlog.Logger) *OrderHandler {
	return &OrderHandler{Cfg: cfg, Log: logger}
}

func (h *OrderHandler) HandleOrdersEvent(msg []byte) {
	var msgObj models.Message
	if err := json.Unmarshal(msg, &msgObj); err != nil {
		log.Println("Error parseando tipo de evento:", err)
		return
	}

	switch msgObj.EventType {
	case "order.created":
		var evt models.OrderEventCreated
		err := json.Unmarshal(msg, &evt)
		if err != nil {
			log.Fatalf("Error parsing %v", err)
			return
		}
		subject := "Nueva orden creada"
		body := BuildOrderCreatedEmailText(evt)
		err = email.SendEmail(h.Cfg, subject, body)
		if err != nil {
			return
		}

	case "order.confirmed":
		var evt models.OrderConfirmedEvent
		err := json.Unmarshal(msg, &evt)
		if err != nil {
			log.Fatalf("Error parsing %v", err)
			return
		}
		subject := "Orden confirmada"
		body := BuildOrderConfirmedEmailText(evt)
		err = email.SendEmail(h.Cfg, subject, body)
		if err != nil {
			return
		}

	case "order.status_updated":
		var evt models.OrderStatusUpdatedEvent
		_ = json.Unmarshal(msg, &evt)
		subject := "Estado de orden actualizado"
		body := fmt.Sprintf("Orden %s cambió de %s a %s", evt.OrderID, evt.PreviousStatus, evt.NewStatus)
		err := email.SendEmail(h.Cfg, subject, body)
		if err != nil {
			return
		}

	default:
		log.Println("⚠️ Tipo de evento no reconocido:", msgObj.EventType)
	}

}
