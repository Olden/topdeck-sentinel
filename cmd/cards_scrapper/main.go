package main

import (
	"fmt"

	"github.com/olden/topdeck-sentinel/internal/sentinel"
	"github.com/olden/topdeck-sentinel/pkg/config"
	"github.com/olden/topdeck-sentinel/pkg/scryfall"
)

func main() {
	sc, err := scryfall.NewScryfallClient()
	if err != nil {
		fmt.Printf("%+v", err)
	}
	cards, err := sc.GetAllCards()
	if err != nil {
		fmt.Printf("%+v", err)
	}

	c, err := config.NewConfig()
	if err != nil {
		fmt.Printf("%+v", err)
	}

	cr, err := sentinel.NewCardsRepo(c.Mysql)
	if err != nil {
		fmt.Printf("%+v", err)
	}

	err = cr.StoreCards(cards)
	if err != nil {
		fmt.Printf("%+v", err)
	}
}
