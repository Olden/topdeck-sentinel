package sentinel

import (
	"github.com/olden/topdeck-sentinel/pkg/config"
	"github.com/olden/topdeck-sentinel/pkg/database"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type LotRepo struct {
	mysql *gorm.DB
}

func NewLotsRepo(c *config.MysqlConfig) (*CardRepo, error) {
	db, err := database.NewMysql(c)
	if err != nil {
		return nil, errors.Wrap(err, "can't instantiate gorm")
	}

	return &CardRepo{
		mysql: db,
	}, nil
}

func (cr *CardRepo) StoreLots(col []Lot, st *StopWord) error {
	for _, l := range col {
		l.NormalizeLot(st)
		recognizedCard, err := cr.FindByName(l.LotNormalized)
		if err == nil {
			l.OracleID = recognizedCard.OracleID
		}
		_ = cr.mysql.FirstOrCreate(&l, Lot{ID: l.ID})
	}

	return nil
}
