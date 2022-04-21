package sentinel

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	AuctionUriPattern = "https://topdeck.ru/apps/toptrade/auctions/%s"
)

type Lot struct {
	ID            string `gorm:"primaryKey"`
	OracleID      string `gorm:"index;size:36"`
	Lot           string `json:"lot"`
	LotNormalized string
	DateEstimated string `json:"date_estimated"`
	CurrentBid    string `json:"current_bid"`
	ImageUrl      string `json:"image_url"`
}

func (l *Lot) Url() string {
	return fmt.Sprintf(AuctionUriPattern, l.ID)
}

func (l *Lot) NormalizeLot(st *StopWord) {
	s := strings.Split(l.Lot, "/")
	tmp := s[0]
	s = strings.Split(tmp, "\\")
	tmp = s[0]

	re := regexp.MustCompile(`\(.+?\)`)
	tmp = re.ReplaceAllString(tmp, "")

	re = regexp.MustCompile(`\[.+?\]`)
	tmp = re.ReplaceAllString(tmp, "")

	re = regexp.MustCompile(`([^\s]+)-([^\s]+)`)
	tmp = re.ReplaceAllString(tmp, "$1 $2")

	re = regexp.MustCompile(`\(.+?\)|[.,\/#№!$%\^&\*;:{}=\-_~()]`)
	tmp = re.ReplaceAllString(tmp, " ")

	re = regexp.MustCompile(`\d+\s*[xх]|\s+[xх]\s*\d`)
	tmp = re.ReplaceAllString(tmp, "")

	tmp = st.CleanString(tmp)

	re = regexp.MustCompile(`\d+`)
	tmp = re.ReplaceAllString(tmp, "")

	re = regexp.MustCompile(`\s\s+`)
	tmp = strings.Trim(re.ReplaceAllString(tmp, " "), " ")
	l.LotNormalized = tmp
}
