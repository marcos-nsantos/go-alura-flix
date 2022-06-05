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

func (vr *VideoRepository) FindAllVideos(pagination *models.Pagination) (*[]models.Video, error) {
	var videos *[]models.Video
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuilder := vr.db.Model(&models.Video{}).Order("created_at desc")
	err := queryBuilder.Count(&pagination.TotalRows).Offset(offset).Limit(pagination.Limit).Find(&videos).Error
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

func (vr *VideoRepository) FindAllVideosByTitle(title string, pagination *models.Pagination) (*[]models.Video, error) {
	var videos *[]models.Video
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuilder := vr.db.Model(&models.Video{}).Where("titulo like ?", "%"+title+"%")
	err := queryBuilder.Count(&pagination.TotalRows).Offset(offset).Limit(pagination.Limit).Find(&videos).Error
	return videos, err
}
