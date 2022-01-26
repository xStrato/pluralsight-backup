package main

import (
	"log"

	"github.com/xStrato/pluralsight-backup/backend/infrastructure/database/contexts"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	bkpDbContext := contexts.NewBackupContext(db)
	bkpDbContext.RunMigrations()

	println("running")
}
