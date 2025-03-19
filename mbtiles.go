//go:build mbtiles

package show

import (
	"context"
	"fmt"
	"os"

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

	fs.Usage = func() {
		fmt.Fprintf(os.Stderr, "Command-line tool for serving MBTiles tiles from an on-demand web server.\n")
		fmt.Fprintf(os.Stderr, "Usage:\n\t%s mbtiles [options]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Valid options are:\n")
		fs.PrintDefaults()
	}

	fs.Parse(args)

	opts, err := mbtiles.RunOptionsFromFlagSet(ctx, fs)

	if err != nil {
		return err
	}

	return mbtiles.RunWithOptions(ctx, opts)
}
