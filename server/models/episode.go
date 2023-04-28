package models

type Episode struct {
	ID               int    `json:"id" gorm:"primary_key:auto_increment"`
	TitleEpisode     string `json:"titleepisode" form:"titleepisode" gorm:"type: varchar(255)"`
	ThumbnailEpisode string `json:"thumbnailepisode" form:"thumbnailepisode" gorm:"type: varchar(255)"`
	LinkFilm         string `json:"linkfilm" gorm:"type: varchar(255)" form:"linkfilm"`
	FilmID           int    `json:"filmid" form:"filmid"`
	Film             Film   `json:"film"`
	// CategoryID    int      `json:"categoryID"`
	// Category      Category `json:"category"`
}

// type EpisodeFilmResponse struct {
// 	ID            int    `json:"id"`
// 	Title         string `json:"title"`
// 	ThumbnailFilm string `json:"thumbnailfilm"`
// 	LinkFilm      string `json:"linkfilm"`
// }

// func (EpisodeFilmResponse) TableName() string {
// 	return "episode"
// }
