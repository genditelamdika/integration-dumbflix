package models

import "time"

type Transaction struct {
	ID        int                  `json:"id" gorm:"primary_key:auto_increment"`
	StartDate time.Time            `json:"startdate" `
	EndDate   time.Time            `json:"enddate" `
	UserID    int                  `json:"user_id"`
	User      UsersProfileResponse `json:"user"`
	// UserID    int                  `json:"userid"`
	// User      User   `json:"user"`
	Price int `json:"price"`
	// Status    string               `json:"status"  gorm:"type:varchar(25)"`
	// Attache string `json:"attache" binding:"required, attache" gorm:"unique; not null" `
	Status string `json:"status" gorm:"type: varchar(255)"`
}

// Films []Film `json:"films"`
