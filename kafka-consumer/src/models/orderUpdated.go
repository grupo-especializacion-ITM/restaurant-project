package models

import (
	"strings"
	"time"
)

type CustomTime struct {
	time.Time
}

const customFormat = "2006-01-02 15:04:05.999999"

type OrderConfirmedEvent struct {
	EventID     string     `json:"event_id"`
	EventType   string     `json:"event_type"`
	Timestamp   CustomTime `json:"timestamp"`
	Version     string     `json:"version"`
	OrderID     string     `json:"order_id"`
	CustomerID  string     `json:"customer_id"`
	TotalAmount float64    `json:"total_amount"`
}

type OrderStatusUpdatedEvent struct {
	EventID        string     `json:"event_id"`
	EventType      string     `json:"event_type"`
	Timestamp      CustomTime `json:"timestamp"`
	Version        string     `json:"version"`
	OrderID        string     `json:"order_id"`
	CustomerID     string     `json:"customer_id"`
	PreviousStatus string     `json:"previous_status"`
	NewStatus      string     `json:"new_status"`
}

func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	t, err := time.Parse(customFormat, s)
	if err != nil {
		return err
	}
	ct.Time = t
	return nil
}
