package commands

import (
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"github.com/webbgeorge/new-relic-insights/newrelic"
)

func GetDashboard(logger *logrus.Logger) func(c *cli.Context) error {
	return func(c *cli.Context) error {
		apiKey := c.GlobalString("apiKey")
		if apiKey == "" {
			return errors.New("missing API key")
		}

		region := c.GlobalString("region")
		if apiKey == "" {
			return errors.New("missing region")
		}

		dashboardId := c.String("dashboardId")
		if dashboardId == "" {
			return errors.New("missing dashboard ID")
		}

		outputPath := c.String("output")
		if outputPath == "" {
			return errors.New("missing output path")
		}

		logger.Infof("Getting dashboard '%s' and saving to '%s' \n", dashboardId, outputPath)

		newRelic, err := newrelic.CreateClient(apiKey, region, logger)
		if err != nil {
			return errors.New(fmt.Sprintf("failed to create new relic client: %+v", err))
		}

		dashboard, err := newRelic.GetDashboard(dashboardId)
		if err != nil {
			return errors.New(fmt.Sprintf("failed to get dashboard '%s': %+v", dashboardId, err))
		}

		err = ioutil.WriteFile(outputPath, dashboard, 0755)
		if err != nil {
			return errors.New(fmt.Sprintf("failed to save dashboard '%s': %+v", dashboardId, err))
		}

		logger.Info("Dashboard saved successfully")
		return nil
	}
}
