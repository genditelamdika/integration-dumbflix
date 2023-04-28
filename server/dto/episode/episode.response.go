package episodedto

type EpisodeResponse struct {
	ID               int    `json:"id" gorm:"primary_key:auto_increment"`
	TitleEpisode     string `json:"titleepisode" gorm:"type: varchar(255)"`
	ThumbnailEpisode string `json:"thumbnailepisode" form:"thumbnailepisode" gorm:"type: varchar(255)"  validate:"required"`
	LinkFilm         string `json:"linkfilm" gorm:"type: varchar(255)"`
	FilmID           int    `json:"filmID"`
	// Film          models.Film `json:"film"`
}
