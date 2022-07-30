package lib

import (
	"github.com/jyury11/skeleton/cmd/lib/di"
)

// CreateArgs Create Skeleton Args
type CreateArgs struct {
	ServiceName string
	Src         string
	Dst         string
	Values      string
	IsForce     bool
}

// Create Create Skeleton
func Create(args CreateArgs) error {
	lib, clean, err := di.InitializeLibrary()
	if err != nil {
		return err
	}
	defer clean()

	return lib.Create(
		args.ServiceName,
		args.Src,
		args.Dst,
		args.Values,
		args.IsForce,
	)
}
