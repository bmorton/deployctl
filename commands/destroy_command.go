package commands

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/codegangsta/cli"
)

// NewDestroyCommand returns the CLI command for "destroy".
func NewDestroyCommand() cli.Command {
	return cli.Command{
		Name:        "destroy",
		ShortName:   "x",
		Usage:       "destroy all instances of a given service and version running on the cluster",
		Description: "destroy [service] [version]",
		Action: func(ctx *cli.Context) {
			service, version, err := extractServiceVersionParameters(ctx)
			if err != nil {
				errorAndBail(err)
			}

			c := getClientFromContext(ctx)
			resp, err := c.DestroyDeploy(service, version)
			if err != nil {
				errorAndBail(err)
			}

			defer resp.Body.Close()

			if resp.StatusCode == 204 {
				fmt.Println("Destroy triggered!")
			} else {
				fmt.Println("Destroy failed :(")
				fmt.Println("-- Response Status:", resp.Status)
				body, _ := ioutil.ReadAll(resp.Body)
				fmt.Println("-- Response Body:", string(body))

				os.Exit(1)
			}
		},
	}
}
