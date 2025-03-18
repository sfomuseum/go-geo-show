//go:build mbtiles
package show

import (
	"context"

	mbtiles "github.com/sfomuseum/go-mbtiles-show"
)

type MbtilesCommand struct {
	Command
}

func init() {
	ctx := context.Background()
	RegisterCommand(ctx, "mbtiles", NewMbtilesCommand)
}

func NewMbtilesCommand(ctx context.Context, cmd string) (Command, error) {
	c := &MbtilesCommand{}
	return c, nil
}

func (c *MbtilesCommand) Run(ctx context.Context, args []string) error {

	fs := mbtiles.DefaultFlagSet()
	fs.Parse(args)

	opts, err := mbtiles.RunOptionsFromFlagSet(ctx, fs)

	if err != nil {
		return err
	}

	return mbtiles.RunWithOptions(ctx, opts)
}
