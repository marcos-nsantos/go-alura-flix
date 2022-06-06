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

func (vr *VideoRepository) FindAllVideos(pagination *models.Pagination[models.Video]) error {
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuilder := vr.db.Model(&models.Video{}).Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)
	return queryBuilder.Find(&pagination.Results).Count(&pagination.TotalRows).Error
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

func (vr *VideoRepository) FindAllVideosByTitle(title string, pagination *models.Pagination[models.Video]) error {
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuilder := vr.db.Model(&models.Video{}).Where("titulo LIKE ?", "%"+title+"%").Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)
	return queryBuilder.Find(&pagination.Results).Count(&pagination.TotalRows).Error
}

func (vr *VideoRepository) FirstThreeVideos(video *[]models.Video) error {
	return vr.db.Limit(3).Find(video).Error
}
