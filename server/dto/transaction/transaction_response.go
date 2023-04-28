package transactiondto

import (
	"dumbmerch/models"
	"time"
)

type TransactionResponse struct {
	ID        int                         `json:"id" gorm:"primary_key:auto_increment"`
	StartDate time.Time                   `json:"startdate" `
	EndDate   time.Time                   `json:"enddate"`
	UserID    int                         `json:"user_id"`
	User      models.UsersProfileResponse `json:"user"`

	Price   int    `json:"price"`
	Attache string `json:"attache" gorm:"type: varchar(255)"`
	Status  string `json:"status" gorm:"type: varchar(255)"`
}

// UserID    string `json:"userid" `
