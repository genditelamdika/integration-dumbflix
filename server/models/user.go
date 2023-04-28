package models

import "time"

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name" gorm:"type: varchar(255)"`
	Email    string `json:"email" gorm:"type: varchar(255)"`
	Password string `json:"-" gorm:"type: varchar(255)"`
	Phone    string `json:"phone" gorm:"type: varchar(255)"`
	Gender   string `json:"gender" gorm:"type: varchar(255)"`
	Address  string `json:"address" gorm:"type: varchar(255)"`
	Role     string `json:"role"`
	Subcribe bool   `json:"subcribe" gorm:"type: bool"`
	// Products  []ProductUserResponse `json:"products"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type UsersProfileResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email" gorm:"type: varchar(255)"`
}

func (UsersProfileResponse) TableName() string {
	return "users"
}

// Profile   ProfileResponse       `json:"profile" binding: "required, email" gorm:"unique;not null"`
// package models

// type Category struct {
// 	ID   int    `json:"id"`
// 	Name string `json:"name" gorm:"type:varchar(255)"`
// 	// Films []Film `json:"films"`
// }

// type CategoryResponse struct {
// 	// ID   int    `json:"id"`
// 	Name string `json:"name"`
// }

// func (CategoryResponse) TableName() string {
// 	return "category"
// }

// package models

// import "time"

// type Film struct {
// 	ID            int    `json:"id"  gorm:"primary_key:auto_increment"`
// 	Title         string `json:"title" gorm:"type: varchar(255)"`
// 	ThumbnailFilm string `json:"thumbnailfilm" gorm:"type: varchar(255)"`
// 	Description   string `json:"description" gorm:"type:text" `
// 	Year          int    `json:"year" gorm:"type: int"`
// 	// Category      CategoryResponse `json:"category"`
// 	// Category_Film []int            `json:"category_film"`
// 	CreatedAt time.Time `json:"-"`
// 	UpdatedAt time.Time `json:"-"`
// }

// type FilmCategoryResponse struct {
// 	ID            int    `json:"id"  gorm:"primary_key:auto_increment"`
// 	Title         string `json:"title" gorm:"type: varchar(255)"`
// 	ThumbnailFilm string `json:"thumbnailfilm"  gorm:"type: varchar(255)"`
// 	Description   string `json:"description" gorm:"type:text" `
// }

// func (FilmCategoryResponse) TableName() string {
// 	return "film"
// }
