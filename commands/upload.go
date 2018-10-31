package commands

import (
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/gaw508/new-relic-insights/newrelic"
	"github.com/urfave/cli"
)

func Upload(c *cli.Context) error {
	apiKey := c.GlobalString("apiKey")
	if apiKey == "" {
		return errors.New("missing API key")
	}

	dashboardId := c.String("dashboardId")
	if dashboardId == "" {
		return errors.New("missing dashboard ID")
	}

	inputPath := c.String("input")
	if inputPath == "" {
		return errors.New("missing input path")
	}

	fmt.Printf("Uploading dashboard '%s' to '%s' \n", dashboardId, inputPath)

	data, err := ioutil.ReadFile(inputPath)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to read file '%s': %+v", inputPath, err))
	}

	newRelic, err := newrelic.CreateClient(apiKey)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to create new relic client: %+v", err))
	}

	err = newRelic.UpdateDashboard(dashboardId, data)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to update dashboard '%s': %+v", dashboardId, err))
	}

	fmt.Println("Complete")
	return nil
}
