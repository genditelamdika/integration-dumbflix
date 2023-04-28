package models

type Category struct {
	ID    int    `json:"id" gorm:"primary_key:auto_increment"`
	Name  string `json:"name" gorm:"type:varchar(255)"`
	Films []Film `json:"film"`
}
type CategoryRelation struct {
	ID            int    `json:"id"`
	Title         string `json:"title"`
	ThumbnailFilm string `json:"thumbnailfilm"`
	Description   string `json:"desc"`
	Year          int    `json:"year"`
}

func (CategoryRelation) tableName() string {
	return "categories"
}
