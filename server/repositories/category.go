package repositories

import (
	"dumbmerch/models"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	FindCategory() ([]models.Category, error)
	GetCategory(ID int) (models.Category, error)
	CreateCategory(category models.Category) (models.Category, error)
	UpdateCategory(category models.Category) (models.Category, error)
	DeleteCategory(category models.Category, ID int) (models.Category, error)
}

func RepositoryCategory(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindCategory() ([]models.Category, error) {
	var categories []models.Category
	err := r.db.Preload("Films").Find(&categories).Error // add this code

	return categories, err
}
func (r *repository) GetCategory(ID int) (models.Category, error) {
	var category models.Category
	err := r.db.Preload("Film").First(&category, ID).Error // add this code

	return category, err
}

func (r *repository) CreateCategory(category models.Category) (models.Category, error) {
	err := r.db.Create(&category).Error

	return category, err
}

func (r *repository) UpdateCategory(category models.Category) (models.Category, error) {
	err := r.db.Save(&category).Error

	return category, err
}
func (r *repository) DeleteCategory(category models.Category, ID int) (models.Category, error) {
	err := r.db.Delete(&category, ID).Scan(&category).Error

	return category, err
}
