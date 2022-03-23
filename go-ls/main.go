package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "all-files",
				Aliases: []string{"a"},
				Usage:   "-a",
			},
			&cli.BoolFlag{
				Name:    "show-user-owner",
				Aliases: []string{"u"},
				Usage:   "-u",
			},
			&cli.BoolFlag{
				Name:    "show-group-owner",
				Aliases: []string{"g"},
				Usage:   "-g",
			},
		},
		Action: func(c *cli.Context) error {
			dirs, files, err := ListFiles(c.Args().First())
			if err != nil {
				return err
			}

			output := NewOutput(dirs, files)
			output.ShowHidden = c.Bool("all-files")
			output.ShowUserOwner = c.Bool("show-user-owner")
			output.ShowGroupOwner = c.Bool("show-group-owner")
			output.Print()

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
