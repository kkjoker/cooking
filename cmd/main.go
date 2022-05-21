package main

import (
	"github.com/kkjoker/recipe/library/rdbms"
	"github.com/kkjoker/recipe/infra"
)

func main() {

	rdbms.Init()
	infra.Init()
}
