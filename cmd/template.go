// ******************************************************************************
// Matronator © 2024.                                                          *
// ******************************************************************************

package cmd

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gookit/color"
	"github.com/urfave/cli/v3"
	"mtrgen/storage"
)

func SaveCommand(_ context.Context, cmd *cli.Command) error {
	name := cmd.Args().First()

	if name == "" {
		return cli.Exit("No name specified.", 1)
	}

	filename := cmd.String("path")

	if filename == "" {
		return cli.Exit("No path specified.", 1)
	}

	store := storage.New()

	err := store.SaveTemplate(name, filename)

	if err != nil {
		color.Errorln("Template doesn't exist on a given path.")
	} else {
		color.Successln("Template saved!")
	}

	return err
}

func RemoveCommand(_ context.Context, cmd *cli.Command) error {
	name := cmd.Args().First()

	store := storage.New()

	err := store.RemoveTemplate(name)

	return err
}

func ListCommand(_ context.Context, cmd *cli.Command) error {
	store := storage.New()

	entries := store.ListEntries()

	i := 1

	green := color.FgLightGreen.Render

	for k := range entries {
		_, err := fmt.Fprintln(cmd.Root().Writer, green(color.Bold.Render(strconv.Itoa(i)+".")), green(k))

		if err != nil {
			color.Errorln(err)
		}

		i++
	}

	return nil
}
