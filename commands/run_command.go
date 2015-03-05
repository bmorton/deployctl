package commands

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/codegangsta/cli"
)

// NewRunCommand returns the CLI command for "run".
func NewRunCommand() cli.Command {
	return cli.Command{
		Name:        "run",
		ShortName:   "r",
		Usage:       "run a task for a given service and version",
		Description: "run [service] [version] [command]",
		Action: func(ctx *cli.Context) {
			service, version, command, err := extractRunParameters(ctx)
			if err != nil {
				errorAndBail(err)
			}

			c := getClientFromContext(ctx)
			resp, err := c.CreateTask(service, version, command)
			if err != nil {
				errorAndBail(err)
			}

			defer resp.Body.Close()

			if resp.StatusCode == 200 {
				io.Copy(os.Stdout, resp.Body)
			} else {
				fmt.Println("Task failed to launch :(")
				fmt.Println("-- Response Status:", resp.Status)
				body, _ := ioutil.ReadAll(resp.Body)
				fmt.Println("-- Response Body:", string(body))

				os.Exit(1)
			}
		},
	}
}

func extractRunParameters(ctx *cli.Context) (service string, version string, command string, err error) {
	if len(ctx.Args()) == 0 {
		return "", "", "", errors.New("Service required")
	}
	service = ctx.Args()[0]

	if len(ctx.Args()) == 1 {
		return service, "", "", errors.New("Version required")
	}
	version = ctx.Args()[1]

	if len(ctx.Args()) == 2 {
		return service, version, "", errors.New("Command required")
	}
	command = strings.Join(ctx.Args()[2:], " ")

	return service, version, command, nil
}
