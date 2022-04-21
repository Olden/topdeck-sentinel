package scryfall

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
	AllCards     = "all_cards"
)

var (
	ErrCardNotFound = errors.New("card not found")
)

type ScryfallClient struct {
	scryfall.Client
}

func NewScryfallClient() (*ScryfallClient, error) {
	client, err := scryfall.NewClient()
	if err != nil {
		return nil, errors.Wrap(err, "can't instantiate scryfall client")
	}

	return &ScryfallClient{
		Client: *client,
	}, nil
}

func (sc *ScryfallClient) GetDefaultCards() ([]sentinel.Card, error) {
	return sc.getCards(DefaultCards)
}

func (sc *ScryfallClient) GetAllCards() ([]sentinel.Card, error) {
	return sc.getCards(AllCards)
}

func (sc *ScryfallClient) FindCardByName(n string) (sentinel.Card, error) {
	ctx := context.Background()

	sco := scryfall.SearchCardsOptions{
		IncludeMultilingual: true,
	}
	result, err := sc.SearchCards(ctx, n, sco)
	if err != nil {
		return sentinel.Card{}, errors.Wrap(err, "can't search card at scryfall")
	}
	if result.TotalCards == 1 {
		return sentinel.Card{
			ID:          result.Cards[0].ID,
			OracleID:    result.Cards[0].OracleID,
			Name:        result.Cards[0].Name,
			Set:         result.Cards[0].Set,
			ScryfallURI: result.Cards[0].ScryfallURI,
		}, nil
	}

	return sentinel.Card{}, ErrCardNotFound
}

func (sc *ScryfallClient) getCards(t string) ([]sentinel.Card, error) {
	ctx := context.Background()

	result, err := sc.ListBulkData(ctx)
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
