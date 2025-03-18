//go:build geojson
package show

import (
	"context"

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
	fs.Parse(args)

	opts, err := geojson.RunOptionsFromFlagSet(ctx, fs)

	if err != nil {
		return err
	}

	return geojson.RunWithOptions(ctx, opts)
}
