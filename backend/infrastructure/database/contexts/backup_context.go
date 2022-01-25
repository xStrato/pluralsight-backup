package contexts

import (
	"log"

	"github.com/xStrato/pluralsight-backup/backend/domain/entities"
	"gorm.io/gorm"
)

type BackupContext struct {
	db *gorm.DB
}

func NewBackupContext(db *gorm.DB) *BackupContext {
	return &BackupContext{db}
}

func (v *BackupContext) RunMigrations() {
	if err := v.db.AutoMigrate(&entities.Video{}); err != nil {
		log.Fatalln("Cannot run migrations for Video entity: ", err.Error())
	}
}

func (v *BackupContext) GetDBConnection() *gorm.DB {
	return v.db
}
