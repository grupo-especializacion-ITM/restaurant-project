package handlers

import (
	"fmt"
	"strings"

	"kafka.consumer.go/src/models"
)

func BuildOrderCreatedEmailText(event models.OrderEventCreated) string {
	var b strings.Builder

	b.WriteString("🧾 NUEVA ORDEN CREADA\n\n")
	b.WriteString(fmt.Sprintf("ID de la Orden: %s\n", event.OrderID))
	b.WriteString(fmt.Sprintf("Cliente: %s\n", event.CustomerID))
	b.WriteString(fmt.Sprintf("Fecha: %s\n", event.Timestamp.Time.Format("02-Jan-2006 15:04:05")))
	b.WriteString(fmt.Sprintf("Estado: %s\n\n", event.Status))
	b.WriteString("🛒 Detalles de los Items:\n")
	b.WriteString("Producto\tCantidad\tPrecio Unit.\tTotal\n")

	for _, item := range event.Items {
		b.WriteString(fmt.Sprintf(
			"%s\t%d\t\t$%.2f\t\t$%.2f\n",
			item.Name, item.Quantity, item.UnitPrice, item.TotalPrice,
		))
	}

	b.WriteString(fmt.Sprintf("\nTOTAL A PAGAR: $%.2f\n", event.TotalAmount))
	b.WriteString("\nGracias por usar nuestro sistema de pedidos.")

	return b.String()
}

func BuildOrderConfirmedEmailText(event models.OrderConfirmedEvent) string {
	var b strings.Builder

	b.WriteString("✅ ORDEN CONFIRMADA\n\n")
	b.WriteString(fmt.Sprintf("ID de la Orden : %s\n", event.OrderID))
	b.WriteString(fmt.Sprintf("Cliente        : %s\n", event.CustomerID))
	b.WriteString(fmt.Sprintf("Fecha          : %s\n", event.Timestamp.Time.Format("02-Jan-2006 15:04:05")))
	b.WriteString(fmt.Sprintf("Versión Evento : %s\n", event.Version))
	b.WriteString(fmt.Sprintf("ID del Evento  : %s\n", event.EventID))
	b.WriteString(fmt.Sprintf("Tipo de Evento : %s\n", event.EventType))
	b.WriteString(fmt.Sprintf("\n💰 Total a Pagar: $%.2f\n", event.TotalAmount))

	b.WriteString("\nGracias por tu compra. Tu orden ha sido confirmada exitosamente.")

	return b.String()
}
