package models

import "time"

type Film struct {
	ID            int      `json:"id" gorm:"primary_key:auto_increment"`
	Title         string   `json:"title" form:"title" gorm:"type: varchar(255)"`
	ThumbnailFilm string   `json:"thumbnailfilm" form:"thumbnailfilm" gorm:"type: varchar(255)"`
	Year          int      `json:"year" form:"year" gorm:"type: int"`
	Linkfilm      string   `json:"linkfilm" form:"linkfilm" gorm:"type: varchar(255)"`
	CategoryID    int      `json:"categoryID" `
	Category      Category `json:"category" `

	// EpisodeID     int                 `json:"episodeID" `
	// Episode       EpisodeFilmResponse `json:"episode" `

	Description string    `json:"description" form:"description" gorm:"type:text" `
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}

type FilmCategoryResponse struct {
	ID            int    `json:"id" form:"id"`
	Title         string `json:"title" form:"title"`
	ThumbnailFilm string `json:"thumbnailfilm" form:"thumbnailfilm"`
	Description   string `json:"description" form:"description"`
	Year          int    `json:"year" form:"year"`
}

func (FilmCategoryResponse) TableName() string {
	return "Film"
}
