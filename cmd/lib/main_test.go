package lib_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/Jyury11/skeleton/cmd/lib"
)

func TestMain(t *testing.T) {
	t.Run("main_lib", func(t *testing.T) {
		p, _ := os.Getwd()
		root := filepath.Join(p, "..", "..")
		src := filepath.Join(root, "example", "template")
		dst := filepath.Join(root, "example", "dst")
		val := filepath.Join(root, "example", "values.yaml")

		args := lib.CreateArgs{
			ServiceName: "user",
			Src:         src,
			Dst:         dst,
			Values:      val,
		}
		if err := lib.Create(args); err != nil {
			panic(err)
		}
	})
}
