package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/olden/topdeck-sentinel/internal/sentinel"
	"github.com/olden/topdeck-sentinel/pkg/config"
	"github.com/olden/topdeck-sentinel/pkg/scryfall"
)

func main() {
	sc, err := scryfall.NewScryfallClient()
	if err != nil {
		fmt.Printf("%+v", err)
	}

	card, err := sc.FindCardByName("Illuna, Apex of Wishes / Иллуна, Венец Желаний")
	spew.Dump(card)
	spew.Dump(err)
	return
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

	cr.FindByOracleId(card.OracleID)
}
