package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	scryfall "github.com/BlueMonday/go-scryfall"
	"github.com/olden/topdeck-sentinel/internal/sentinel"
	"github.com/pkg/errors"
)

const (
	DefaultCards = "default_cards"
)

type CardScrapper struct {
	scryfall *scryfall.Client
}

func NewCardScrapper() (*CardScrapper, error) {
	client, err := scryfall.NewClient()
	if err != nil {
		return nil, errors.Wrap(err, "can't instantiate scryfall client")
	}

	return &CardScrapper{
		scryfall: client,
	}, nil
}

func (cs *CardScrapper) getCards(t string) ([]sentinel.Card, error) {
	ctx := context.Background()

	result, err := cs.scryfall.ListBulkData(ctx)
	if err != nil {
		fmt.Printf("%+v", err)
	}

	var url string
	for _, item := range result {
		if item.Type == t {
			url = item.DownloadURI
		}
	}
	if url == "" {
		return nil, errors.New("can't get bulk file for cards")
	}

	resp, err := downloadFile(url)
	if err != nil {
		return nil, errors.Wrap(err, "can't download bulk file")
	}

	var cards []sentinel.Card
	err = json.Unmarshal(resp, &cards)
	if err != nil {
		return nil, errors.Wrap(err, "can't unmarshal cards data from bulk file")
	}

	return cards, nil
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
