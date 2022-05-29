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

func (vr *VideoRepository) FindByID(id uint) (*models.Video, error) {
	var video models.Video
	err := vr.db.First(&video, id).Error
	return &video, err
}

func (vr *VideoRepository) Update(video *models.Video) error {
	return vr.db.Save(video).Error
}
