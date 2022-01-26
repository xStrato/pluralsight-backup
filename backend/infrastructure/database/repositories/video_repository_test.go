package repositories_test

import (
	"fmt"
	"log"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"github.com/xStrato/pluralsight-backup/backend/domain/entities"
	"github.com/xStrato/pluralsight-backup/backend/infrastructure/database/contexts"
	"github.com/xStrato/pluralsight-backup/backend/infrastructure/database/repositories"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestVideoRepositoryMethods(t *testing.T) {
	// db, err := gorm.Open(sqlite.Open("app_test.db"), &gorm.Config{})
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	encoderContext := contexts.NewBackupContext(db)
	encoderContext.RunMigrations()
	videoRepo := repositories.NewVideoRepository(encoderContext)

	var insertedVideoUUID string

	t.Run("Insert_ValidVideoStruct_ShouldNotReturnError", func(t *testing.T) {
		//Arrange
		video := entities.NewVideoWithParams(
			uuid.NewV4().String(),
			uuid.NewV4().String(),
			uuid.NewV4().String(),
			"ModuleTitle",
			"Title",
			"Author",
			"Description",
			"CourseName",
			"Level",
			"Image",
			"",
			"1",
		)
		insertedVideoUUID = video.GetId()
		//Act
		entity, err := videoRepo.Insert(video)
		//Assert
		require.NotNil(t, entity)
		require.Nil(t, err)
		require.Equal(t, video.ID, entity.GetId())
		require.IsType(t, video, entity)
	})

	t.Run("Find_ValidInsertedVideoUUID_ShouldNotReturnError", func(t *testing.T) {
		//Arrange
		//Act
		entity, err := videoRepo.Find(insertedVideoUUID)
		//Assert
		require.NotNil(t, entity)
		require.Nil(t, err)
		require.Equal(t, insertedVideoUUID, entity.GetId())
	})

	t.Run("Find_InvalidInsertedVideoUUID_ShouldReturnError", func(t *testing.T) {
		//Arrange
		invalidID := uuid.NewV4().String()
		errMsg := fmt.Sprintf("video ID '%v' doest not exist", invalidID)
		//Act
		entity, err := videoRepo.Find(invalidID)
		//Assert
		require.Nil(t, entity)
		require.EqualError(t, err, errMsg)
	})
}
