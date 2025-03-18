//go:build geoparquet
package show

import (
	"context"

	geoparquet "github.com/sfomuseum/go-geoparquet-show"
)

type GeoparquetCommand struct {
	Command
}

func init() {
	ctx := context.Background()
	RegisterCommand(ctx, "geoparquet", NewGeoparquetCommand)
}

func NewGeoparquetCommand(ctx context.Context, cmd string) (Command, error) {
	c := &GeoparquetCommand{}
	return c, nil
}

func (c *GeoparquetCommand) Run(ctx context.Context, args []string) error {

	fs := geoparquet.DefaultFlagSet()
	fs.Parse(args)

	opts, err := geoparquet.RunOptionsFromFlagSet(ctx, fs)

	if err != nil {
		return err
	}

	return geoparquet.RunWithOptions(ctx, opts)
}
