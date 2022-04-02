package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/olden/topdeck-sentinel/internal/sentinel"
	"github.com/olden/topdeck-sentinel/pkg/config"
)

func main() {
	resp, err := http.Get(TDApiUri)
	if err != nil {
		fmt.Printf("%+v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%+v", err)
	}

	var lots []sentinel.Lot
	if err := json.Unmarshal(body, &lots); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	c, err := config.NewConfig()
	if err != nil {
		fmt.Printf("%+v", err)
	}

	lr, err := sentinel.NewLotsRepo(c.Mysql)
	if err != nil {
		fmt.Printf("%+v", err)
	}
	err = lr.StoreLots(lots)
	if err != nil {
		fmt.Printf("%+v", err)
	}
}
