package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/olden/topdeck/internal/sentinel"
	"github.com/olden/topdeck/pkg/config"
	"github.com/olden/topdeck/pkg/database"
	"github.com/pkg/errors"
	"gorm.io/gorm/clause"

	scryfall "github.com/BlueMonday/go-scryfall"
)

const (
	BulkType = "default_cards"
)

func main() {
	ctx := context.Background()
	client, err := scryfall.NewClient()
	if err != nil {
		fmt.Printf("%+v", err)
	}

	result, err := client.ListBulkData(ctx)
	if err != nil {
		fmt.Printf("%+v", err)
	}

	var url string
	for _, item := range result {
		if item.Type == BulkType {
			url = item.DownloadURI
		}
	}

	resp, err := downloadFile(url)
	if err != nil {
		fmt.Printf("%+v", err)
	}

	var cards []sentinel.Card
	err = json.Unmarshal(resp, &cards)
	if err != nil {
		fmt.Printf("%+v", err)
	}

	err = storeCards(cards)
	if err != nil {
		fmt.Printf("%+v", err)
	}
}

func storeCards(col []sentinel.Card) error {
	c, err := config.NewConfig()
	if err != nil {
		return errors.Wrap(err, "can't instantiate config")
	}

	db, err := database.NewMysql(c.Mysql)
	if err != nil {
		return errors.Wrap(err, "can't instantiate gorm")
	}

	db.Clauses(clause.Insert{Modifier: "IGNORE"}).CreateInBatches(col, 5000)

	return nil
}

func downloadFile(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad status: %s", resp.Status)
	}

	return io.ReadAll(resp.Body)
}
