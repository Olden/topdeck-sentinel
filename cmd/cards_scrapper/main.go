package main

import (
	"fmt"

	"github.com/olden/topdeck-sentinel/internal/sentinel"
	"github.com/olden/topdeck-sentinel/pkg/config"
)

func main() {
	cs, err := NewCardScrapper()
	if err != nil {
		fmt.Printf("%+v", err)
	}
	cards, err := cs.getCards(DefaultCards)
	if err != nil {
		fmt.Printf("%+v", err)
	}

	c, err := config.NewConfig()
	if err != nil {
		fmt.Printf("%+v", err)
	}

	cr, err := sentinel.NewCardRepo(c.Mysql)
	if err != nil {
		fmt.Printf("%+v", err)
	}

	cr.StoreCards(cards)
	if err != nil {
		fmt.Printf("%+v", err)
	}
}
