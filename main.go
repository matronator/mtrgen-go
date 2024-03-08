package main

import (
	"context"
	"log"
	"os"

	"github.com/urfave/cli/v3"
	"mtrgen/cmd"
)

func main() {
	command := &cli.Command{
		Name:    "mtrgen",
		Usage:   "MTRGen is a tool to generate code from a given template.",
		Version: "1.0",
		Action: func(_ context.Context, cmd *cli.Command) error {
			cli.VersionPrinter(cmd)

			return nil
		},
		Copyright: "Matronator © 2024",
		Commands:  cmd.Commands,
	}

	if err := command.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
