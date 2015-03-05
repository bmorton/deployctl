package main

import (
	"os"

	"github.com/bmorton/deployctl/commands"
	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "deployctl"
	app.Version = version
	app.Usage = "A simple command line client for deployster."
	app.Author = "Brian Morton"
	app.Email = "brian@mmm.hm"
	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "url", Value: "http://localhost:3000", Usage: "a URL to the deployster instance", EnvVar: "DEPLOYSTER_URL"},
		cli.StringFlag{Name: "username", Value: "deployster", Usage: "username for authenticating to deployster", EnvVar: "DEPLOYSTER_USERNAME"},
		cli.StringFlag{Name: "password", Value: "", Usage: "password for authenticating to deployster", EnvVar: "DEPLOYSTER_PASSWORD"},
	}
	app.Commands = []cli.Command{
		commands.NewDeployCommand(),
		commands.NewRunCommand(),
		commands.NewDestroyCommand(),
		commands.NewListCommand(),
	}
	app.Run(os.Args)
}
