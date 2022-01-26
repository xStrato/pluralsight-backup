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
	Level         string `json:"level" valid:"required"`
	Image         string `json:"image" valid:"url"`
	Ordering      string `json:"ordering" valid:"required"`
	Url           string `json:"url" valid:"url"`
	StoragePath   string `json:"storagePath" valid:"-"`
}

func NewVideoWithParams(courseId, moduleId, clipId, moduleTitle, title, authorName, desc, courseName, level, image, url, ordering string) *Video {
	return &Video{
		Entity:      *common.NewEntity(),
		CourseId:    courseId,
		ModuleId:    moduleId,
		ClipId:      clipId,
		ModuleTitle: moduleTitle,
		Title:       title,
		Author:      authorName,
		Description: desc,
		CourseName:  courseName,
		Level:       level,
		Image:       image,
		Url:         url,
		Ordering:    ordering,
	}
}

func NewVideo() *Video {
	return &Video{Entity: *common.NewEntity()}
}

func (v *Video) GetId() string {
	return v.ID
}
