package repository

import (
	"github.com/marcos-nsantos/alura-flix/models"
	"gorm.io/gorm"
)

type VideoRepository struct {
	db *gorm.DB
}

func NewVideoRepository(db *gorm.DB) *VideoRepository {
	return &VideoRepository{db: db}
}

func (vr *VideoRepository) Create(video *models.Video) error {
	return vr.db.Create(video).Error
}

func (vr *VideoRepository) FindAll() ([]*models.Video, error) {
	var videos []*models.Video
	err := vr.db.Find(&videos).Error
	return videos, err
}
