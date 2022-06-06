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

func (vr *VideoRepository) FindAllVideos(videos *[]models.Video) error {
	return vr.db.Find(videos).Error
}

func (vr *VideoRepository) FindVideoByID(video *models.Video) error {
	return vr.db.Where("id = ?", video.ID).First(video).Error
}

func (vr *VideoRepository) UpdateVideo(video *models.Video, videoDataToUpdate *models.Video) error {
	return vr.db.Model(video).Updates(videoDataToUpdate).Error
}

func (vr *VideoRepository) DeleteVideo(video *models.Video) error {
	return vr.db.Delete(video).Error
}

func (vr *VideoRepository) FindAllVideosByTitle(title string, videos *[]models.Video) error {
	return vr.db.Where("titulo LIKE ?", "%"+title+"%").Find(videos).Error
}
