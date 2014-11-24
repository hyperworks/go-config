package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperworks/go-config"
)

type Config struct {
	Env      string
	DBDriver string `config:"DRIVER"`
	DBURL    string `config:"DB"`
}

func main() {
	c := &Config{}
	e := config.ReadEnv(c, "CFG_")
	if e != nil {
		panic(e)
	}

	bytes, _ := json.MarshalIndent(c, "", "  ")
	fmt.Println(string(bytes))
}
