package repositories

import (
	"fmt"

	"github.com/xStrato/pluralsight-backup/backend/domain/entities"
	"github.com/xStrato/pluralsight-backup/backend/domain/interfaces"
	"github.com/xStrato/pluralsight-backup/backend/infrastructure/database/contexts"
)

type VideoRepository struct {
	dbContext *contexts.BackupContext
}

func NewVideoRepository(ctx *contexts.BackupContext) *VideoRepository {
	return &VideoRepository{ctx}
}

func (r *VideoRepository) Insert(e interfaces.Entity) (interfaces.Entity, error) {
	model, ok := e.(*entities.Video)
	if !ok {
		return nil, fmt.Errorf("cannot cast '%v' as Video entity", e)
	}
	db := r.dbContext.GetDBConnection()
	if err := db.Create(&model).Error; err != nil {
		return nil, err
	}
	return model, nil
}

func (r *VideoRepository) Find(id string) (interfaces.Entity, error) {
	var video entities.Video
	db := r.dbContext.GetDBConnection()
	db.First(&video, "id=?", id)

	if video.IsValid() != nil {
		return nil, fmt.Errorf("video ID '%v' doest not exist", id)
	}
	return &video, nil
}

func (r *VideoRepository) Update(e interfaces.Entity) (interfaces.Entity, error) {
	return nil, nil
}
