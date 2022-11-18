package main

import (
	"go-notebook/demo-go-kit/model"
	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "../../../model/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	dsn := "host=localhost user=postgres password=password dbname=demo port=5432 sslmode=disable"
	gormdb, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	g.UseDB(gormdb) // reuse your gorm db

	// Generate basic type-safe DAO API for struct following conventions
	g.ApplyBasic(model.Article{})

	// Generate the code
	g.Execute()
}
