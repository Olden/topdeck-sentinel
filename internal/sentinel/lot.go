package sentinel

import (
	"fmt"

	"gorm.io/gorm"
)

const (
	AuctionUriPattern = "https://topdeck.ru/apps/toptrade/auctions/%s"
)

type Lot struct {
	ID            string `gorm:"primaryKey"`
	DateEstimated string `json:"date_estimated"`
	CurrentBid    string `json:"current_bid"`
	ImageUrl      string `json:"image_url"`
}

func (l *Lot) Url() string {
	return fmt.Sprintf(AuctionUriPattern, l.ID)
}

func (l *Lot) AfterCreate(tx *gorm.DB) (err error) {
	// fire event to notify users
	fmt.Println("lot create: " + l.Url())
	return
}
