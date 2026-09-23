package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "healthChecker",
		Usage: "a simple to check status of websites",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "domain",
				Aliases:  []string{"d"},
				Usage:    "domain name to check",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "port",
				Usage:    "Port number to check",
				Aliases:  []string{"p"},
				Value:    "80",
				Required: false,
			},
		},
		Action: func(ctx *cli.Context) error {
			status := Check(ctx.String("domain"), ctx.String("port"))
			fmt.Println(status)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
