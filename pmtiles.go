//go:build pmtiles

package show

import (
	"context"
	"fmt"
	"os"

	pmtiles "github.com/sfomuseum/go-pmtiles-show"
)

type PMTilesCommand struct {
	Command
}

func init() {
	ctx := context.Background()
	RegisterCommand(ctx, "pmtiles", NewPMTilesCommand)
}

func NewPMTilesCommand(ctx context.Context, cmd string) (Command, error) {
	c := &PMTilesCommand{}
	return c, nil
}

func (c *PMTilesCommand) Run(ctx context.Context, args []string) error {

	fs := pmtiles.DefaultFlagSet()

	fs.Usage = func() {
		fmt.Fprintf(os.Stderr, "Command-line tool for serving PMTiles tiles from an on-demand web server.\n")
		fmt.Fprintf(os.Stderr, "Usage:\n\t%s pmtiles [options]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Valid options are:\n")
		fs.PrintDefaults()
	}

	fs.Parse(args)

	opts, err := pmtiles.RunOptionsFromFlagSet(ctx, fs)

	if err != nil {
		return err
	}

	return pmtiles.RunWithOptions(ctx, opts)
}
