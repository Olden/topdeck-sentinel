package sentinel

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/olden/topdeck-sentinel/pkg/config"
	"github.com/olden/topdeck-sentinel/pkg/database"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func (cr *CardRepo) FindByOracleId(id string) (Card, error) {
	var card Card
	cr.mysql.Where(&Card{OracleID: id}).First(&card)
	spew.Dump(card)

	return card, nil
}
