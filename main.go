package main

import (
	"fmt"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "isodate",
		Usage: "Show the current time in RFC3339",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "time_zone",
				Aliases: []string{"tz"},
			},
		},
		Action: func(c *cli.Context) error {
			now := time.Now()
			tzString := c.String("time_zone")
			if tzString != "" {
				tz, err := time.LoadLocation("America/New_York")
				if err != nil {
					return err
				}
				now = now.In(tz)
			}

			fmt.Println(now.Format(time.RFC3339))
			return nil
		},
		EnableBashCompletion: true,
	}

	app.Run(os.Args)
}
