package main

import (
	"os"

	"github.com/codegangsta/cli"
	cfg "github.com/ziyadparekh/runtime/config"
)

func main() {
	app := cli.NewApp()
	app.Name = cfg.APP_NAME
	app.Usage = cfg.APP_USAGE
	app.Version = cfg.APP_VERSION
	app.Author = cfg.APP_AUTHOR
	app.Email = cfg.APP_EMAIL
	app.Commands = cfg.APP_COMMANDS

	app.Run(os.Args)
}
