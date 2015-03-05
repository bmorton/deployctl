package commands

import (
	"errors"

	"github.com/codegangsta/cli"
)

func extractServiceVersionParameters(ctx *cli.Context) (service string, version string, err error) {
	if len(ctx.Args()) == 0 {
		return "", "", errors.New("Service required")
	}
	service = ctx.Args()[0]

	if len(ctx.Args()) == 1 {
		return service, "", errors.New("Version required")
	}
	version = ctx.Args()[1]

	return service, version, nil
}
