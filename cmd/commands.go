// ******************************************************************************
// Matronator © 2024.                                                          *
// ******************************************************************************

package cmd

import (
	"github.com/urfave/cli/v3"
)

var Commands = []*cli.Command{
	{
		Name:    "add",
		Aliases: []string{"a"},
		Usage:   "Get template from online registry and add it to the local store.",
	},
	{
		Name:    "generate",
		Aliases: []string{"gen", "g"},
		Usage:   "Generates a file from template. The template can be specified by name if it's saved in the local store, or with a path to the template file relative to the CWD.",
	},
	{
		Name:    "login",
		Aliases: []string{"in"},
		Usage:   "Login to the online registry.",
	},
	{
		Name:    "parse",
		Aliases: []string{"p"},
		Usage:   "Parses a regular file without a template header using the parser.",
	},
	{
		Name:    "profile",
		Aliases: []string{"whoami"},
		Usage:   "Shows the current logged user if any.",
		Action:  ProfileCommand,
	},
	{
		Name:    "publish",
		Aliases: []string{"pub"},
		Usage:   "Publish template to the online template repository.",
	},
	{
		Name:    "remove",
		Aliases: []string{"rm"},
		Usage:   "Removes a template from the local store.",
		Action:  RemoveCommand,
	},
	{
		Name:    "repair",
		Aliases: []string{"r"},
		Usage:   "Repairs the local store if broken or missing templates.",
	},
	{
		Name:    "save",
		Aliases: []string{"s"},
		Usage:   "Saves a template from file to the local store.",
		Action:  SaveCommand,
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "path", Usage: "Path to template", Aliases: []string{"p"}, Required: true, Persistent: false},
		},
	},
	{
		Name:    "save-bundle",
		Aliases: []string{"sb", "saveb", "save-b"},
		Usage:   "Creates a bundle from two or more template files and add it to your local store.",
	},
	{
		Name:    "show",
		Aliases: []string{"ls"},
		Usage:   "List all saved templates in the local store.",
		Action:  ListCommand,
	},
	{
		Name:    "signup",
		Aliases: []string{"sign-up", "su", "up", "sign"},
		Usage:   "Create new user account in the online registry.",
	},
	{
		Name:    "use",
		Aliases: []string{"u"},
		Usage:   "Generates a file using a template from the online registry without saving it.",
	},
	{
		Name:    "validate",
		Aliases: []string{"v"},
		Usage:   "Check if a file is valid template or bundle.",
	},
}
