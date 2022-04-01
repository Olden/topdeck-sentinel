package sentinel

import "github.com/google/uuid"

type Card struct {
	ID          uuid.UUID `json:"id", gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name        string    `json:"name"`
	Set         string    `json:"set"`
	ScryfallURI string    `json:"scryfall_uri"`
}
