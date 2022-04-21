package sentinel

type Card struct {
	ID          string `json:"id" gorm:"primaryKey;size:36"`
	OracleID    string `json:"oracle_id" gorm:"index;size:36;not null"`
	Name        string `json:"name" gorm:"index:idx_card_names,class:FULLTEXT"`
	PrintedName string `json:"printed_name" gorm:"index:idx_card_names,class:FULLTEXT"`
	Language    string `json:"lang"`
	Set         string `json:"set"`
	ScryfallURI string `json:"scryfall_uri"`
}
