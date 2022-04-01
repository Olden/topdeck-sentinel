package main

import (
	"fmt"

	"github.com/olden/topdeck/internal/sentinel"
	"github.com/olden/topdeck/pkg/config"
	"github.com/olden/topdeck/pkg/database"
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
