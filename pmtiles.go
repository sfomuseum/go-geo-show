//go:build pmtiles
package show

import (
	"context"

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
	fs.Parse(args)

	opts, err := pmtiles.RunOptionsFromFlagSet(ctx, fs)

	if err != nil {
		return err
	}

	return pmtiles.RunWithOptions(ctx, opts)
}
