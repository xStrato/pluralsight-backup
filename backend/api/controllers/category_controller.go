package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xStrato/pluralsight-backup/backend/application/common"
	"github.com/xStrato/pluralsight-backup/backend/application/services"
	"github.com/xStrato/pluralsight-backup/backend/domain/entities"
)

type CategoryController struct {
	service *services.VideoService
}

func NewCategoryController(service *services.VideoService) *CategoryController {
	return &CategoryController{
		service: service,
	}
}

func (ctr *CategoryController) Create(ctx *gin.Context) {
	model := entities.NewVideo()
	if err := ctx.ShouldBindJSON(&model); err != nil {
		ctx.JSON(http.StatusBadRequest, common.NewGenericResult(false, "Cannot bind JSON from body", err.Error()))
		return
	}
	result, err := ctr.service.Insert(model)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, common.NewGenericResult(false, "Insert operation failed", err.Error()))
		return
	}
	ctx.JSONP(http.StatusCreated, result)
}
