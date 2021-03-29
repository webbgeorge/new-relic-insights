package commands

import (
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"github.com/webbgeorge/new-relic-insights/newrelic"
)

func CreateDashboard(logger *logrus.Logger) func(c *cli.Context) error {
	return func(c *cli.Context) error {
		apiKey := c.GlobalString("apiKey")
		if apiKey == "" {
			return errors.New("missing API key")
		}

		region := c.GlobalString("region")
		if apiKey == "" {
			return errors.New("missing region")
		}

		inputPath := c.String("input")
		if inputPath == "" {
			return errors.New("missing input path")
		}

		logger.Infof("Creating dashboard '%s' \n", inputPath)

		data, err := ioutil.ReadFile(inputPath)
		if err != nil {
			return errors.New(fmt.Sprintf("failed to read file '%s': %+v", inputPath, err))
		}

		newRelic, err := newrelic.CreateClient(apiKey, region, logger)
		if err != nil {
			return errors.New(fmt.Sprintf("failed to create new relic client: %+v", err))
		}

		err = newRelic.CreateDashboard(data)
		if err != nil {
			return errors.New(fmt.Sprintf("failed to create dashboard: %+v", err))
		}

		logger.Info("Dashboard created successfully")
		return nil
	}
}
