package kafka

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	kitlog "github.com/go-kit/log"
)

// Logger global
var Logger kitlog.Logger

func Consume(ctx context.Context, topic string, handler func([]byte)) error {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":  "localhost:9092",
		"group.id":           "go-consumer-group",
		"auto.offset.reset":  "earliest",
		"enable.auto.commit": false,
	})
	if err != nil {
		return err
	}
	defer c.Close()

	err = c.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		return err
	}

	slog.Info("Kafka consumer started", "topic", topic)

	run := true
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	for run {
		select {
		case <-ctx.Done():
			run = false
		case sig := <-sigchan:
			slog.Info("caught signal", "signal", sig)
			run = false
		default:
			ev := c.Poll(100)
			if ev == nil {
				continue
			}

			switch msg := ev.(type) {
			case *kafka.Message:
				slog.Info("Message received",
					"topic", *msg.TopicPartition.Topic,
					"offset", msg.TopicPartition.Offset,
					"message", string(msg.Value),
				)

				handler(msg.Value)

				_, err := c.CommitMessage(msg)
				if err != nil {
					slog.Error("Failed to commit message", "err", err)
				}
			case kafka.Error:
				slog.Error("Kafka error", "code", msg.Code(), "err", msg)
			default:
				// other events are ignored
			}
		}
	}

	slog.Info("Kafka consumer stopped")
	return nil
}

/* func Consume(ctx context.Context, topic string, handler func([]byte)) {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "go-prod-group",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		level.Error(Logger).Log("msg", "failed to create consumer", "error", err)
		return
	}

	consumer.SubscribeTopics([]string{topic}, nil)
	level.Info(Logger).Log("msg", "subscribed", "topic", topic)

	for {
		select {
		case <-ctx.Done():
			_ = consumer.Close()
			return
		default:
			msg, err := consumer.ReadMessage(100 * time.Millisecond)
			if err != nil {
				continue // no bloquear por timeouts
			}
			handler(msg.Value)
		}
	}
} */
