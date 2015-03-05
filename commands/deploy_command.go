package commands

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/codegangsta/cli"
)

// NewDeployCommand returns the CLI command for "deploy".
func NewDeployCommand() cli.Command {
	return cli.Command{
		Name:        "deploy",
		ShortName:   "d",
		Usage:       "Deploys a given service and version to the cluster",
		Description: "deploy [service] [version]",
		Flags: []cli.Flag{
			cli.BoolFlag{Name: "skip-destroy-previous", Usage: "skip old version destruction after successful deploy (if possible)"},
			cli.IntFlag{Name: "instances", Usage: "number of instances to deploy (optional, defaults to 0, which means current number running or 1 if none running)"},
		},
		Action: func(ctx *cli.Context) {
			service, version, err := extractServiceVersionParameters(ctx)
			if err != nil {
				errorAndBail(err)
			}

			if ctx.Bool("skip-destroy-previous") {
				fmt.Println("Skipping destruction of previous version...")
			}

			c := getClientFromContext(ctx)
			resp, err := c.CreateDeploy(service, version, !ctx.Bool("skip-destroy-previous"), ctx.Int("instances"))
			if err != nil {
				errorAndBail(err)
			}

			defer resp.Body.Close()

			if resp.StatusCode == 201 {
				fmt.Println("Deploy triggered!")
			} else {
				fmt.Println("Deploy failed :(")
				fmt.Println("-- Response Status:", resp.Status)
				body, _ := ioutil.ReadAll(resp.Body)
				fmt.Println("-- Response Body:", string(body))

				os.Exit(1)
			}
		},
	}
}
