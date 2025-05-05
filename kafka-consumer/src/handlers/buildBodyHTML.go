package handlers

import (
	"fmt"
	"strings"

	"kafka.consumer.go/src/models"
)

func BuildOrderCreatedEmailBody(event models.OrderEventCreated) string {
	var b strings.Builder

	b.WriteString(fmt.Sprintf("🧾 <b>Nueva Orden Creada</b><br/><br/>"))
	b.WriteString(fmt.Sprintf("<b>ID de la Orden:</b> %s<br/>", event.OrderID))
	b.WriteString(fmt.Sprintf("<b>Cliente:</b> %s<br/>", event.CustomerID))
	b.WriteString(fmt.Sprintf("<b>Fecha:</b> %s<br/>", event.Timestamp.Time.Format("02-Jan-2006 15:04:05")))
	b.WriteString(fmt.Sprintf("<b>Estado:</b> %s<br/>", event.Status))
	b.WriteString("<br/><b>🛒 Detalles de los Items:</b><br/>")
	b.WriteString("<table border='1' cellpadding='5' cellspacing='0'>")
	b.WriteString("<tr><th>Producto</th><th>Cantidad</th><th>Precio Unitario</th><th>Total</th></tr>")

	for _, item := range event.Items {
		b.WriteString(fmt.Sprintf(
			"<tr><td>%s</td><td>%d</td><td>$%.2f</td><td>$%.2f</td></tr>",
			item.Name, item.Quantity, item.UnitPrice, item.TotalPrice,
		))
	}

	b.WriteString("</table><br/>")
	b.WriteString(fmt.Sprintf("<b>Total a pagar:</b> <span style='color:green;'>$%.2f</span><br/>", event.TotalAmount))
	b.WriteString("<br/>Gracias por usar nuestro sistema de pedidos.")

	return b.String()
}
