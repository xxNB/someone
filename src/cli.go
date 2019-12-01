package main

import (
	"log"
	"os"
	"fmt"

	"github.com/urfave/cli/v2"
)

func cliLow() {
	var doption string

	app := &cli.App{
		Flags: []cli.Flag {
			&cli.StringFlag{
				Name:        "doption",
				Value:       "top250",
				Usage:       "get newmoive from douban",
				Destination: &doption,
			},
		},
		Action: func(c *cli.Context) error {
			name := "someone"
			if c.NArg() > 0 {
				name = c.Args().Get(0)
			}
			if doption == "newMovie" {
				ss := DouMovie{}
				ss.newHignMovieSelect()
			} else {
				fmt.Println("Hello", name)
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}