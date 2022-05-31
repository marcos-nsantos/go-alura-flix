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

func (vr *VideoRepository) CreateVideo(video *models.Video) error {
	return vr.db.Create(video).Error
}

func (vr *VideoRepository) FindAllVideos() (*[]models.Video, error) {
	var videos *[]models.Video
	err := vr.db.Find(&videos).Error
	return videos, err
}

func (vr *VideoRepository) FindVideoByID(id uint) (*models.Video, error) {
	var video models.Video
	err := vr.db.First(&video, id).Error
	return &video, err
}

func (vr *VideoRepository) UpdateVideo(video *models.Video) error {
	return vr.db.Save(video).Error
}

func (vr *VideoRepository) DeleteVideo(video *models.Video) error {
	return vr.db.Delete(video).Error
}

func (vr *VideoRepository) FindAllVideosByTitle(title string) (*[]models.Video, error) {
	var videos *[]models.Video
	err := vr.db.Where("titulo LIKE ?", "%"+title+"%").Find(&videos).Error
	return videos, err
}
