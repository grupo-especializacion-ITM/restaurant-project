package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	kitlog "github.com/go-kit/log"

	"github.com/go-kit/log/level"
	"kafka.consumer.go/src/config"
	"kafka.consumer.go/src/handlers"
	"kafka.consumer.go/src/kafka"
)

// Logger global
var Logger kitlog.Logger

func main() {

	Logger = kitlog.NewLogfmtLogger(os.Stdout)
	Logger = level.NewFilter(Logger, level.AllowAll())

	level.Info(Logger).Log("Initialize")

	ctx, cancel := context.WithCancel(context.Background())

	conf := config.InitEnv()

	orderHandler := handlers.NewOrderHandler(conf, &Logger)

	// Wrap orderHandler to match the expected function signature
	orderHandlerFunc := func(msg []byte) {
		orderHandler.HandleOrdersEvent(msg)
	}

	go kafka.Consume(ctx, conf.OrderTopic, orderHandlerFunc)
	/* go kafka.Consume(ctx, conf.InventoryTopic, func(msg []byte) {
		log.Println("📄 Log recibido:", string(msg))
	}) */

	// Shutdown graceful
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	level.Info(Logger).Log("msg", "shutdown signal received")
	cancel()
	time.Sleep(time.Second * 1) // da tiempo a cerrar
}
