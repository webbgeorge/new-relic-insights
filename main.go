package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"github.com/webbgeorge/new-relic-insights/cli_logger"
	"github.com/webbgeorge/new-relic-insights/commands"
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
		cli.StringFlag{
			Name:   "region, r",
			Usage:  "New Relic Insights API region",
			EnvVar: "NEW_RELIC_INSIGHTS_API_REGION",
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
			Name:        "dashboard",
			Usage:       "options for dashboards",
			Subcommands: []cli.Command{
				{
					Name:    "get",
					Usage:   "get dashboard as JSON file",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "dashboardId, d",
							Usage: "id of the dashboard",
						},
						cli.StringFlag{
							Name:  "output, o",
							Usage: "path of output file for dashboard JSON",
						},
					},
					Action: commands.GetDashboard(logger),
				},
				{
					Name:    "create",
					Usage:   "create a new dashboard from JSON file",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "input, i",
							Usage: "path of input file for dashboard JSON",
						},
					},
					Action: commands.CreateDashboard(logger),
				},
				{
					Name:    "update",
					Usage:   "update dashboard from JSON file",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "dashboardId, d",
							Usage: "id of the dashboard to update",
						},
						cli.StringFlag{
							Name:  "input, i",
							Usage: "path of input file for dashboard JSON",
						},
					},
					Action: commands.UpdateDashboard(logger),
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		logger.Fatal(err)
	}
}
