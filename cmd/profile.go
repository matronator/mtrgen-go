// ******************************************************************************
// Matronator © 2024.                                                          *
// ******************************************************************************

package cmd

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
	"mtrgen/registry"
)

func ProfileCommand(_ context.Context, _ *cli.Command) error {
	profile := registry.New()

	if profile.Profile.Username != "" {
		fmt.Println("Currently logged-in as:", profile.Profile.Username)
	}

	fmt.Println("No user logged in.")

	return nil
}
