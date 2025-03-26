package main

import (
	"log"
	"mirea_finance_tracker/internal/config"
	"mirea_finance_tracker/internal/model"
	"mirea_finance_tracker/internal/router"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
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

	r := router.SetupRouter(db)
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
