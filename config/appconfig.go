package config

import (
	"fmt"
	"log"
	"net/http"

	"github.com/codegangsta/cli"
	"github.com/ziyadparekh/runtime/repository"
	"github.com/ziyadparekh/runtime/router"
)

const (
	APP_NAME    = "protobox"
	APP_USAGE   = "Configure your app runtime programmaticaly"
	APP_VERSION = "0.1"
	APP_AUTHOR  = "Ziyad Parekh"
	APP_EMAIL   = "ziyad@1stdibs.com"
)

var APP_COMMANDS = []cli.Command{
	{
		Name:  "init",
		Usage: "Initialize your runtime description",
		Action: func(c *cli.Context) {
			fmt.Println("====> Initializing new runtime configuration")
			r := repository.NewRuntime()
			fmt.Println(string(r))
		},
	},
	{
		Name:  "setup",
		Usage: "Setup your machine once the config is written",
		Action: func(c *cli.Context) {
			fmt.Println("====> Creating directories and all that fun stuff")
			repository.SetupStructure()
		},
	},
	{
		Name:  "install",
		Usage: "Save a dependency to your runtime description",
		Action: func(c *cli.Context) {
			if len(c.Args()) == 0 {
				fmt.Println("No package name provided")
				return
			}
			fmt.Println("====> Attempting to save the following dependencies:")
			if _, err := repository.SaveDependencies(c.Args()); err == nil {
				fmt.Println("====> Runtime description updated")
			}
		},
	},
	{
		Name:  "run",
		Usage: "Run a server listening on port 8000",
		Action: func(c *cli.Context) {
			fmt.Println("====> Starting server on port 8000")
			router := router.NewRouter()
			log.Fatal(http.ListenAndServe(":8000", router))

		},
	},
}
