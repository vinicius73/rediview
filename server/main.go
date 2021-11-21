package main

import (
	"fmt"

	"github.com/vinicius73/rediview/pkg/config"
)

func main() {
	cfg := config.Build()

	fmt.Println(cfg)
}
