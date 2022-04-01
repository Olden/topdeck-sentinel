package main

import (
	"fmt"

	"github.com/olden/topdeck-sentinel/internal/sentinel"
	"github.com/olden/topdeck-sentinel/pkg/config"
	"github.com/olden/topdeck-sentinel/pkg/database"
)

func main() {
	c, err := config.NewConfig()
	if err != nil {
		fmt.Printf("%+v", err)
	}

	db, err := database.NewMysql(c.Mysql)
	if err != nil {
		fmt.Printf("%+v", err)
	}

	err = db.AutoMigrate(&sentinel.Card{})
	if err != nil {
		fmt.Printf("%+v", err)
	}
}
