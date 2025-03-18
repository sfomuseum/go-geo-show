package show

import (
	"context"
	"fmt"
	"os"

	"github.com/sfomuseum/go-geo-show"
)

func usage() {

	fmt.Println("Usage: show [CMD] [OPTIONS]")
	fmt.Println("Valid commands are:")

	for _, cmd := range show.Commands() {
		fmt.Printf("* %s\n", cmd)
	}

	os.Exit(0)
}

func Run(ctx context.Context) error {

	if len(os.Args) < 2 {
		usage()
	}

	cmd := os.Args[1]

	if cmd == "-h" {
		usage()
	}

	c, err := show.NewCommand(ctx, cmd)

	if err != nil {
		usage()
	}

	args := make([]string, 0)

	if len(os.Args) > 2 {
		args = os.Args[2:]
	}

	err = c.Run(ctx, args)

	if err != nil {
		return fmt.Errorf("Failed to run '%s' command, %w", cmd, err)
	}

	return nil
}
