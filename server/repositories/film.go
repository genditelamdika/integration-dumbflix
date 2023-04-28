package repositories

import (
	"dumbmerch/models"

	"gorm.io/gorm"
)

type FilmRepository interface {
	FindFilm() ([]models.Film, error)
	GetFilm(ID int) (models.Film, error)
	CreateFilm(film models.Film) (models.Film, error)
	UpdateFilm(film models.Film, Id int) (models.Film, error)
	DeleteFilm(film models.Film) (models.Film, error)
	GetCategoryfilm(ID int) (models.Category, error)
}

func RepositoryFilm(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindFilm() ([]models.Film, error) {
	var films []models.Film
	err := r.db.Preload("Category").Find(&films).Error

	return films, err
}

func (r *repository) GetFilm(ID int) (models.Film, error) {
	var film models.Film
	err := r.db.Preload("Category").First(&film, ID).Error

	return film, err
}

func (r *repository) CreateFilm(film models.Film) (models.Film, error) {
	err := r.db.Create(&film).Error

	return film, err
}

func (r *repository) UpdateFilm(film models.Film, Id int) (models.Film, error) {
	err := r.db.Model(&film).Updates(&film).Error

	return film, err
}

func (r *repository) DeleteFilm(film models.Film) (models.Film, error) {
	err := r.db.Delete(&film).Error

	return film, err
}
func (r *repository) GetCategoryfilm(Id int) (models.Category, error) {
	var cate models.Category
	err := r.db.First(&cate, Id).Error
	return cate, err
	// err := r.db.Delete(&film).Error

	// return cate, err
}
