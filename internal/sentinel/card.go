package sentinel

type Card struct {
	ID          string `json:"id" gorm:"primaryKey;size:36"`
	OracleID    string `json:"oracle_id" gorm:"size:36;uniqueIndex;not null"`
	Name        string `json:"name"`
	Set         string `json:"set"`
	ScryfallURI string `json:"scryfall_uri"`
}
