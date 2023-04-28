package episodedto

import "dumbmerch/models"

type CreateEpisodeRequest struct {
	ID               int         `json:"id" gorm:"primary_key:auto_increment"`
	TitleEpisode     string      `json:"titleepisode" form:"titleepisode" gorm:"type: varchar(255)" validate:"required"`
	ThumbnailEpisode string      `json:"thumbnailepisode" form:"thumbnailepisode" gorm:"type: varchar(255)" validate:"required"`
	LinkFilm         string      `json:"linkfilm" form:"linkfilm" gorm:"type: varchar(255)"`
	FilmID           int         `json:"filmID" form:"filmid"`
	Film             models.Film `json:"film" form:"film" validate:"required"`
	// CategoryID    int             `json:"categoryID"`
	// Category      models.Category `json:"category"`
	// Film          Film   `json:"film
}

type UpdateEpisodeRequest struct {
	TitleEpisode     string      `json:"titleepisode" form:"title" gorm:"type: varchar(255)"`
	ThumbnailEpisode string      `json:"thumbnailepisode" form:"thumbnailepisode" gorm:"type: varchar(255)"`
	LinkFilm         string      `json:"linkfilm" form:"linkfilm" gorm:"type: varchar(255)"`
	FilmID           int         `json:"filmID" form:"filmid"`
	Film             models.Film `json:"film" form:"category" validate:"required"`
}
