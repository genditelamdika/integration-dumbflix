package filmsdto

import "time"

type FilmResponse struct {
	ID            int    `json:"id"`
	Title         string `json:"title" form:"title" gorm:"type: varchar(255)"  validate:"required"`
	ThumbnailFilm string `json:"thumbnailfilm" form:"thumbnailfilm" gorm:"type: varchar(255)"  validate:"required"`
	Description   string `json:"description" gorm:"type:text" form:"desc"  validate:"required"`
	Year          int    `json:"year" gorm:"type: int"  validate:"required"`
	// Linkfilm      string `json:"linkfilm" form:"linkfilm" gorm:"type: varchar(255)"`
	CategoryID int `json:"categoryID"`
	// Category   Category `json:"category"`
	// Episode       []EpisodeFilmResponse `json:"episode"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
