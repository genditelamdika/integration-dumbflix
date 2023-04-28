package filmsdto

import "dumbmerch/models"

type CreateFilmRequest struct {
	Title            string          `json:"title" form:"title" gorm:"type: varchar(255)"`
	ThumbnailFilm    string          `json:"thumbnailfilm" form:"thumbnailfilm" gorm:"type: varchar(255)"`
	Year             int             `json:"year" gorm:"type: int"`
	Description      string          `json:"description" gorm:"type:text" form:"desc"`
	CategoryID       int             `json:"categoryID" form:"category_id"`
	Category         models.Category `json:"category" form:"category" validate:"required"`
	TitleEpisode     string          `json:"title_episode" form:"title_episode" gorm:"type: varchar(255)"`
	ThumbnailEpisode string          `json:"thumbnailepisode" form:"thumbnailepisode" gorm:"type: varchar(255)"`
	Linkfilm         string          `json:"linkfilm" form:"linkfilm" gorm:"type: varchar(255)"`
}

type UpdateFilmRequest struct {
	Title         string `json:"title" form:"title"`
	ThumbnailFilm string `json:"thumbnailfilm" form:"thumbnailfilm"`
	Year          int    `json:"year"`
	// Linkfilm      string          `json:"linkfilm" form:"linkfilm" gorm:"type: varchar(255)"`
	CategoryID  int             `json:"categoryID"`
	Category    models.Category `json:"category" form:"category" validate:"required"`
	Description string          `json:"description"  form:"desc"`
}
