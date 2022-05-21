package main

import (
	"github.com/kkjoker/cooking/library/rdbms"
	"github.com/kkjoker/cooking/infra"
)

func main() {

	rdbms.Init()
	infra.Init()
}
