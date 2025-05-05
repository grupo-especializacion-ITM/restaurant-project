package config

import (
	"fmt"
	"os"
)

type Conf struct {
	EmailFrom      string
	EmailTo        string
	EmailPassword  string
	OrderTopic     string
	InventoryTopic string
	EmailApiKey    string
	EmailSecretKey string
}

var Config *Conf

func InitEnv() *Conf {

	emailFrom := getEnvOrPanic("EMAIL_FROM")
	emailTo := getEnvOrPanic("EMAIL_TO")
	emailPass := getEnvOrPanic("EMAIL_PASSWORD")
	emailApiKey := getEnvOrPanic("EMAIL_API_KEY")
	emailSecretKey := getEnvOrPanic("EMAIL_SECRET_KEY")
	orderTopic := getEnvOrDefault("ORDER_TOPIC", "restaurant.orders")
	inventoryTopic := getEnvOrDefault("INVENTORY_TOPIC", "restaurant.inventory")

	return &Conf{
		EmailFrom:      emailFrom,
		EmailTo:        emailTo,
		EmailPassword:  emailPass,
		OrderTopic:     orderTopic,
		InventoryTopic: inventoryTopic,
		EmailApiKey:    emailApiKey,
		EmailSecretKey: emailSecretKey,
	}
}

/* func validateError(err error) {
	if err != nil {
		panic(err)
	}
} */

func getEnvOrPanic(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok || len(value) == 0 {
		panic(fmt.Sprintf("Environment variable %s not found", key))
	}
	return value
}

func getEnvOrDefault(key string, def string) string {
	value, ok := os.LookupEnv(key)
	if !ok || len(value) == 0 {
		return def
	}
	return value
}
