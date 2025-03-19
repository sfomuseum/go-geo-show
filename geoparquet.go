//go:build geoparquet

package show

import (
	"context"
	"fmt"
	"os"

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

	fs.Usage = func() {
		fmt.Fprintf(os.Stderr, "Command-line tool for serving GeoParquet features as vector tiles from an on-demand web server.\n")
		fmt.Fprintf(os.Stderr, "Usage:\n\t %s geoparquet [options]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Valid options are:\n")
		fs.PrintDefaults()
	}

	fs.Parse(args)

	opts, err := geoparquet.RunOptionsFromFlagSet(ctx, fs)

	if err != nil {
		return err
	}

	return geoparquet.RunWithOptions(ctx, opts)
}
