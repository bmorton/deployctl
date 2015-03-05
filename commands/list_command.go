package commands

import (
	"errors"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/bmorton/deployctl/client"
	"github.com/codegangsta/cli"
)

// NewListCommand returns the CLI command for "list".
func NewListCommand() cli.Command {
	return cli.Command{
		Name:        "list",
		ShortName:   "l",
		Usage:       "list all instances of a given service running on the cluster",
		Description: "list [service]",
		Action: func(ctx *cli.Context) {
			if len(ctx.Args()) == 0 {
				errorAndBail(errors.New("Service required"))
			}
			service := ctx.Args()[0]

			c := getClientFromContext(ctx)
			units, err := c.GetUnits(service)
			if err != nil {
				errorAndBail(err)
			}

			printUnits(units)
		},
	}
}

func printUnits(units []*client.Unit) {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 4, ' ', 0)
	fmt.Fprintln(w, "#\tVersion\tCurrent state\tDesired state\tMachine ID\tTimestamp")
	fmt.Fprintln(w, "-\t-------\t-------------\t-------------\t----------\t---------")
	for _, u := range units {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\n", u.Instance, u.Version, u.CurrentState, u.DesiredState, u.MachineID, u.Timestamp)
	}
	w.Flush()
}
