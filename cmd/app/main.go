package main

import (
	"Programa4/pkg"
	"errors"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	// CI CD
	app := &cli.App{
		Name:  "Tree app",
		Usage: "see a tree and other options",
		Commands: []*cli.Command{
			{
				Name:        "dir",
				Aliases:     []string{"d"},
				Description: "List directories",
				Action: func(context *cli.Context) error {
					if context.NArg() == 0 || context.NArg() > 1 {
						return errors.New("need path")
					}
					if err := pkg.ListDirectories(context.Args().Get(0)); err != nil {
						log.Println("Failed to list directories")
						return err
					}
					return nil
				},
			},
			{
				Name:        "file",
				Aliases:     []string{"f"},
				Description: "List files",
				Action: func(context *cli.Context) error {
					if context.NArg() == 0 || context.NArg() > 1 {
						return errors.New("need path")
					}
					if err := pkg.ListFiles(context.Args().Get(0)); err != nil {
						log.Println("Failed to list files")
						return err
					}
					return nil
				},
			},
		},
		Action: func(context *cli.Context) error {
			var err error
			if context.NArg() == 0 {
				err = pkg.Tree("./", 0)
			} else {
				err = pkg.Tree(context.Args().Get(0), 0)
			}
			if err != nil {
				return err
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
