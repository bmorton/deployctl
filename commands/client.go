package commands

import (
	"github.com/bmorton/deployctl/client"
	"github.com/codegangsta/cli"
)

func getClientFromContext(ctx *cli.Context) *client.Client {
	username := ctx.GlobalString("username")
	password := ctx.GlobalString("password")
	baseURL := ctx.GlobalString("url")
	return client.New(baseURL, username, password)
}
