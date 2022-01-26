package services

import (
	"github.com/xStrato/pluralsight-backup/backend/application/common"
	"github.com/xStrato/pluralsight-backup/backend/domain/interfaces"
	"github.com/xStrato/pluralsight-backup/backend/infrastructure/database/repositories"
)

type VideoService struct {
	repo *repositories.VideoRepository
}

func NewVideoService(repo *repositories.VideoRepository) *VideoService {
	return &VideoService{repo}
}

func (r *VideoService) Insert(e interfaces.Entity) (*common.GenericResult, error) {
	err := e.IsValid()
	if err != nil {
		return nil, err
	}
	result, err := r.repo.Insert(e)
	if err != nil {
		return nil, err
	}
	return common.NewGenericResult(true, "Video was successfully inserted", result), nil
}
