package main

import (
	"os"

	"github.com/gaw508/new-relic-insights/commands"
	"github.com/urfave/cli"
	"github.com/sirupsen/logrus"
	"github.com/gaw508/new-relic-insights/cli_logger"
)

func main() {
	app := cli.NewApp()
	app.Name = "nrinsights"
	app.HelpName = "nrinsights"
	app.Usage = "New Relic Insights CLI"
	app.Version = "0.1.0"
	cli.VersionFlag = cli.BoolFlag{
		Name: "version",
		Usage: "print only the version",
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "apiKey, k",
			Usage:  "New Relic Insights API key",
			EnvVar: "NEW_RELIC_INSIGHTS_API_KEY",
		},
		cli.BoolFlag{
			Name:  "verbose, v",
			Usage: "Verbose output",
		},
	}

	logger := logrus.New()
	logger.SetFormatter(cli_logger.Formatter{})
	app.Before = func(c *cli.Context) error {
		if c.Bool("verbose") {
			logger.SetLevel(logrus.DebugLevel)
		}
		return nil
	}

	app.Commands = []cli.Command{
		{
			Name:    "download",
			Aliases: []string{"dl"},
			Usage:   "download a New Relic Insights dashboard as JSON",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "dashboardId, d",
					Usage: "id of the dashboard to download",
				},
				cli.StringFlag{
					Name:  "output, o",
					Usage: "path of output file for dashboard JSON",
				},
			},
			Action: commands.Download(logger),
		},
		{
			Name:    "upload",
			Aliases: []string{"ul"},
			Usage:   "update a New Relic Insights dashboard using JSON",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "dashboardId, d",
					Usage: "id of the dashboard to download",
				},
				cli.StringFlag{
					Name:  "input, i",
					Usage: "path of input file for dashboard JSON",
				},
			},
			Action: commands.Upload(logger),
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		logger.Fatal(err)
	}
}
