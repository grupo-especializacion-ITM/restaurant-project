package models

type OrderEventCreated struct {
	EventID     string     `json:"event_id"`
	EventType   string     `json:"event_type"`
	Timestamp   CustomTime `json:"timestamp"`
	Version     string     `json:"version"`
	OrderID     string     `json:"order_id"`
	CustomerID  string     `json:"customer_id"`
	Items       []Item     `json:"items"`
	TotalAmount float64    `json:"total_amount"`
	Status      string     `json:"status"`
}

type Item struct {
	ID         string  `json:"id"`
	ProductID  string  `json:"product_id"`
	Name       string  `json:"name"`
	Quantity   int     `json:"quantity"`
	UnitPrice  float64 `json:"unit_price"`
	TotalPrice float64 `json:"total_price"`
}

/* func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	// Remover comillas del string JSON
	s := strings.Trim(string(b), "\"")

	layout := "2006-01-02 15:04:05.000000"
	t, err := time.Parse(layout, s)
	if err != nil {
		return fmt.Errorf("error parsing time: %w", err)
	}

	ct.Time = t
	return nil
} */
