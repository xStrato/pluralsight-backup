package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xStrato/pluralsight-backup/backend/application/common"
	"github.com/xStrato/pluralsight-backup/backend/domain/entities"
	"github.com/xStrato/pluralsight-backup/backend/infrastructure/database/repositories"
)

type CategoryController struct {
	repo *repositories.VideoRepository
}

func NewCategoryController(repo *repositories.VideoRepository) *CategoryController {
	return &CategoryController{
		repo: repo,
	}
}

func (ctr *CategoryController) Create(ctx *gin.Context) {
	model := entities.NewVideo()
	if err := ctx.ShouldBindJSON(&model); err != nil {
		ctx.JSON(http.StatusBadRequest, common.NewGenericResult(false, "Cannot bind JSON from body", err.Error()))
		return
	}
	err := model.IsValid()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, common.NewGenericResult(false, "Request model is not valid", err.Error()))
		return
	}
	result, err := ctr.repo.Insert(model)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, common.NewGenericResult(false, "Insert operation failed", err.Error()))
		return
	}
	ctx.JSONP(http.StatusCreated, common.NewGenericResult(true, "Video was successfully inserted", result))
}
