package sentinel

import (
	"github.com/olden/topdeck-sentinel/pkg/config"
	"github.com/olden/topdeck-sentinel/pkg/database"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CardRepo struct {
	mysql *gorm.DB
}

func NewCardRepo(c *config.MysqlConfig) (*CardRepo, error) {
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
