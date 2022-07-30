package main_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/jyury11/skeleton/cmd/skeleton/di"
)

func TestMain(t *testing.T) {
	t.Run("main", func(t *testing.T) {
		cli, clean, err := di.InitializeCLI()
		if err != nil {
			panic(err)
		}
		defer clean()

		p, _ := os.Getwd()
		root := filepath.Join(p, "..", "..")
		src := filepath.Join(root, "example", "template")
		dst := filepath.Join(root, "example", "dst")
		val := filepath.Join(root, "example", "values.yaml")
		cli.SetArgs([]string{"create", "-s", "user", "--src", src, "--dst", dst, "--values", val})
		if err := cli.Execute(); err != nil {
			panic(err)
		}
	})
}

func TestMainForce(t *testing.T) {
	t.Run("main_force", func(t *testing.T) {
		cli, clean, err := di.InitializeCLI()
		if err != nil {
			panic(err)
		}
		defer clean()

		p, _ := os.Getwd()
		root := filepath.Join(p, "..", "..")
		src := filepath.Join(root, "example", "template")
		dst := filepath.Join(root, "example", "dst")
		val := filepath.Join(root, "example", "values.yaml")
		cli.SetArgs([]string{"create", "-s", "user", "--src", src, "--dst", dst, "--values", val, "-f"})
		if err := cli.Execute(); err != nil {
			panic(err)
		}
	})
}
