package main

import (
	"log"

	"github.com/xStrato/pluralsight-backup/backend/api"
	"github.com/xStrato/pluralsight-backup/backend/api/controllers"
	"github.com/xStrato/pluralsight-backup/backend/api/routes"
	"github.com/xStrato/pluralsight-backup/backend/infrastructure/database/contexts"
	"github.com/xStrato/pluralsight-backup/backend/infrastructure/database/repositories"
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

	videoRepository := repositories.NewVideoRepository(bkpDbContext)
	categoryController := controllers.NewCategoryController(videoRepository)
	categoryRouter := routes.NewCategoryRouter(categoryController)

	server := api.NewServer("3000")
	server.Router(categoryRouter.Configure)

	server.Run()
}
