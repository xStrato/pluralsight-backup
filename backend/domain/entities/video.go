package entities

import (
	"github.com/asaskevich/govalidator"
	"github.com/xStrato/pluralsight-backup/backend/domain/common"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Video struct {
	common.Entity `valid:"required"`
	CourseId      string `json:"courseId" valid:"uuid"`
	ModuleId      string `json:"moduleId" valid:"uuid"`
	ClipId        string `json:"clipId" valid:"uuid"`
	ModuleTitle   string `json:"moduleTitle" valid:"required"`
	Title         string `json:"title" valid:"required"`
	Author        string `json:"authorName" valid:"required"`
	Description   string `json:"description" valid:"required"`
	CourseName    string `json:"courseName" valid:"required"`
	Image         string `json:"image" valid:"url"`
	Ordering      int    `json:"ordering" valid:"required"`
	FilePath      string `json:"filePath" valid:"-"`
}

func NewVideo(rID, file string) *Video {
	return &Video{
		Entity: *common.NewEntity(),
	}
}

func (v *Video) GetId() string {
	return v.ID
}
