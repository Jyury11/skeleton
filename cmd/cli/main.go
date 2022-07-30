package main

import (
	"github.com/jyury11/skeleton/cmd/cli/di"
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
