package sentinel

import (
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/olden/topdeck-sentinel/pkg/config"
	"github.com/olden/topdeck-sentinel/pkg/database"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var (
	ErrNotFound = errors.New("not found")
)

type CardRepo struct {
	mysql *gorm.DB
}

func NewCardsRepo(c *config.MysqlConfig) (*CardRepo, error) {
	db, err := database.NewMysql(c)
	if err != nil {
		return nil, errors.Wrap(err, "can't instantiate gorm")
	}

	return &CardRepo{
		mysql: db,
	}, nil
}

func (cr *CardRepo) StoreCards(col []Card) error {
	cr.mysql.Clauses(clause.Insert{Modifier: "IGNORE"}).CreateInBatches(col, 5000)

	return nil
}

func (cr *CardRepo) GetAll() ([]Card, error) {
	var cards []Card
	result := cr.mysql.Find(&cards)

	if result.Error != nil {
		return cards, result.Error
	}
	if result.RowsAffected > 0 {
		return cards, nil
	} else {
		return cards, ErrNotFound
	}
}

func (cr *CardRepo) FindByOracleId(id string) (Card, error) {
	var card Card
	cr.mysql.Where(&Card{OracleID: id}).First(&card)
	spew.Dump(card)

	return card, nil
}

func (cr *CardRepo) FindByName(n string) (Card, error) {
	var card Card
	var processed []string
	name := strings.Split(n, " ")
	for _, word := range name {
		processed = append(processed, "+"+word)
	}
	binding := strings.Join(processed, " ")
	// fmt.Printf("SELECT * FROM `sentinel`.`cards` where match(name,printed_name) AGAINST ('%s' in boolean mode);\n", binding)
	cr.mysql.Raw("SELECT * FROM `sentinel`.`cards` where match(name,printed_name) AGAINST (? in boolean mode);", binding).Scan(&card)
	if card.OracleID != "" {
		return card, nil
	}

	return card, ErrNotFound
}
