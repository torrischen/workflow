package flow

import (
	"fmt"

	"github.com/torrischen/workflow/logging"
	"github.com/torrischen/workflow/pkg/util"
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

	db.Callback().Create().Before("gorm:create").Register("add:uuid", addUUID)
}

func addUUID(db *gorm.DB) {
	if _, ok := db.Statement.Schema.FieldsByName["ID"]; ok {
		db.Statement.SetColumn("ID", util.Uuid())
	}
}

type Base struct {
	ID        string `gorm:"primaryKey;"`
	CreatedAt int64
	UpdatedAt int64
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
