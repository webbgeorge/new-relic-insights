package commands

import (
	"errors"
	"fmt"

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

	fmt.Printf("Downloading dashboard '%s' to '%s' \n", dashboardId, inputPath)

	// TODO

	fmt.Println("Complete")
	return nil
}
