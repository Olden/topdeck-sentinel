package database

import (
	"fmt"
	"net/url"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/olden/topdeck-sentinel/pkg/config"
	"github.com/pkg/errors"
)

func NewMysql(c *config.MysqlConfig) (*gorm.DB, error) {
	db, err := gorm.Open(
		mysql.Open(fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=%s",
			c.User,
			c.Passwd,
			c.Host,
			c.Port,
			c.DB,
			url.QueryEscape(c.Loc),
		)),
		&gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
		},
	)
	if err != nil {
		errors.Wrap(err, "failed to connect database")
	}

	if err == nil {
		d, err := db.DB()
		if err != nil {
			errors.Wrap(err, "failed to set connection options")
		}
		d.SetMaxOpenConns(c.MaxOpenConns)
		d.SetMaxIdleConns(c.MaxIdleConns)
		d.SetConnMaxLifetime(time.Duration(c.ConnMaxLifetime) * time.Minute)
	}

	return db, err
}
