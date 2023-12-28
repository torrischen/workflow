package flow

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/torrischen/workflow/logging"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

func initMysql(user, password, host, port, dbname string) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, dbname,
	)
	_db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		logging.Fatalf("Init db failed: %s", err)
	}

	db = _db

	// Migrate the schema
	db.AutoMigrate(&Pipeline{})
	db.AutoMigrate(&Node{})
	db.AutoMigrate(&NodeData{})
	db.AutoMigrate(&PipelineRun{})
}

type Base struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	CreatedAt int64
	UpdatedAt int64
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
