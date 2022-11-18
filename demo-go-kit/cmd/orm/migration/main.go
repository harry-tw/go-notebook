package main

import (
	"go-notebook/demo-go-kit/migration"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {
	dsn := "host=localhost user=postgres password=password dbname=demo port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Panic(err)
	}
	if err = migration.Execute(db); err != nil {
		log.Panicf("failed to migrate db: %v", err)
	}
	return
}
