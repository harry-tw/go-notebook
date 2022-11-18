package migration

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

func Execute(db *gorm.DB) error {
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "202211181220",
			Migrate: func(tx *gorm.DB) error {
				type Article struct {
					gorm.Model
					ID      uuid.UUID
					Title   string
					Content string
				}
				return tx.AutoMigrate(&Article{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("articles")
			},
		},
		{
			ID: "202211181810",
			Migrate: func(tx *gorm.DB) error {
				type Article struct {
					CreatedAt time.Time
					UpdatedAt time.Time
					DeletedAt time.Time
				}
				return tx.AutoMigrate(&Article{})
			},
			Rollback: func(tx *gorm.DB) error {
				err := tx.Migrator().DropColumn("articles", "created_at")
				if err != nil {
					return err
				}
				err = tx.Migrator().DropColumn("articles", "updated_at")
				if err != nil {
					return err
				}
				err = tx.Migrator().DropColumn("articles", "deleted_at")
				if err != nil {
					return err
				}
				return err
			},
		},
		// Add migrations above, refer to https://github.com/go-gormigrate/gormigrate
		// To make migration being effective, run `go run main.go` within cmd/orm/migration.
	})
	return m.Migrate()
}
