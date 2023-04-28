package categoriesdto

type CategoryResponse struct {
	ID   int    `json:"id" gorm:"primary_key:auto_increment"`
	Name string `json:"name"`
	// Films []CategoryFilm `json:"films"`
}
type CategoryFilm struct {
	ID            int    `json:"id"`
	Title         string `json:"title"`
	ThumbnailFilm string `json:"thumbnailfilm"`
	Description   string `json:"desc"`
	Year          int    `json:"year"`
}

// type CategoryResponse struct {
// 	ID   int    `json:"id" gorm:"primary_key:auto_increment"`
// 	Name string `json:"name"`
// 	// FilmID int`json:"films"`
// 	// Films []Film `json:"films"`
// }

// type CategoryFilm struct {
// 	ID            int    `json:"id"`
// 	Title         string `json:"title"`
// 	ThumbnailFilm string `json:"thumbnailfilm"`
// 	Description   string `json:"desc"`
// 	Year          int    `json:"year"`
// }

// type CategoryResponse struct {
// 	ID   int    `json:"id"`
// 	Name string `json:"name" form:"name" validate:"required"`
// 	// Films []CategoryFilm `json:"films"`
// 	// Email    string `json:"email" form:"email" validate:"required"`
// 	// Password string `json:"password" form:"password" validate:"required"`
// 	// Phone    string `json:"phone" form:"phone" validate:"required"`
// 	// Gender   string `json:"gender" form:"gender" validate:"required"`
// 	// Address  string `json:"address" form:"address" validate:"required"`
// 	// Subcribe bool   `json:"subcribe" form:"subcribe"`
// }
