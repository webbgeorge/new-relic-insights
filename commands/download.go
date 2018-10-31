package commands

import (
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/gaw508/new-relic-insights/newrelic"
	"github.com/urfave/cli"
)

func Download(c *cli.Context) error {
	apiKey := c.GlobalString("apiKey")
	if apiKey == "" {
		return errors.New("missing API key")
	}

	dashboardId := c.String("dashboardId")
	if dashboardId == "" {
		return errors.New("missing dashboard ID")
	}

	outputPath := c.String("output")
	if outputPath == "" {
		return errors.New("missing output path")
	}

	fmt.Printf("Downloading dashboard '%s' to '%s' \n", dashboardId, outputPath)

	newRelic, err := newrelic.CreateClient(apiKey)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to create new relic client: %+v", err))
	}

	dashboard, err := newRelic.GetDashboard(dashboardId)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to get dashboard '%s': %+v", dashboardId, err))
	}

	ioutil.WriteFile(outputPath, dashboard, 0755)

	fmt.Println("Complete")
	return nil
}
