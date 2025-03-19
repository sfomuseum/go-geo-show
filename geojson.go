//go:build geojson

package show

import (
	"context"
	"fmt"
	"os"

	geojson "github.com/sfomuseum/go-geojson-show"
)

type GeojsonCommand struct {
	Command
}

func init() {
	ctx := context.Background()
	RegisterCommand(ctx, "geojson", NewGeojsonCommand)
}

func NewGeojsonCommand(ctx context.Context, cmd string) (Command, error) {
	c := &GeojsonCommand{}
	return c, nil
}

func (c *GeojsonCommand) Run(ctx context.Context, args []string) error {

	fs := geojson.DefaultFlagSet()

	fs.Usage = func() {
		fmt.Fprintf(os.Stderr, "Command-line tool for serving GeoJSON features from an on-demand web server.\n")
		fmt.Fprintf(os.Stderr, "Usage:\n\t %s geojson path(N) path(N)\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Valid options are:\n")
		fs.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nIf the only path as input is \"-\" then data will be read from STDIN.\n\n")
	}

	fs.Parse(args)

	opts, err := geojson.RunOptionsFromFlagSet(ctx, fs)

	if err != nil {
		return err
	}

	return geojson.RunWithOptions(ctx, opts)
}
