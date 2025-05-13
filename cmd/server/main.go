// @title Mirea Finance Tracker API
// @version 1.0
// @description API для отслеживания личных финансов
// @host localhost:8080
// @BasePath /
// @schemes http

package main

import (
	"log"
	"mirea_finance_tracker/internal/config"
	"mirea_finance_tracker/internal/model"
	"mirea_finance_tracker/internal/redis"
	"mirea_finance_tracker/internal/router"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/sirupsen/logrus"
)

func main() {
	redis.InitRedis()
	logFile, err := os.OpenFile("/app/logs/api.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logrus.SetOutput(logFile)
	} else {
		logrus.Warn("Не удалось создать файл логов, используем stdout")
	}

	cfg := config.Load()

	db, err := gorm.Open(postgres.Open(cfg.DSN()), &gorm.Config{})
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}

	err = db.AutoMigrate(
		&model.User{},
		&model.Currency{},
		&model.Account{},
		&model.Category{},
		&model.Transaction{},
		&model.RecurringTransaction{},
	)
	if err != nil {
		log.Fatalf("Ошибка миграции: %v", err)
	}

	log.Println("Подключение к БД успешно, миграция завершена")

	seedCurrencies(db)

	r := router.SetupRouter(db)
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}

func seedCurrencies(db *gorm.DB) {
	currencies := []model.Currency{
		{Code: "RUB", Name: "Российский рубль"},
		{Code: "USD", Name: "Доллар США"},
		{Code: "EUR", Name: "Евро"},
	}

	for _, currency := range currencies {
		db.FirstOrCreate(&currency, model.Currency{Code: currency.Code})
	}

	log.Println("Базовые валюты инициализированы")
}
