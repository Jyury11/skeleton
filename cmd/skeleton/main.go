package main

import (
	"github.com/jyury11/skeleton/cmd/skeleton/di"
)

func main() {
	cli, clean, err := di.InitializeCLI()
	if err != nil {
		panic(err)
	}
	defer clean()
	if err := cli.Execute(); err != nil {
		panic(err)
	}
}
