package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "rb-ldap"
	app.Usage = "Command line interface for Redbrick LDAP"
	app.Authors = []cli.Author{
		{
			Name:  "Cian Butler",
			Email: "butlerx@redbrick.dcu.ie",
		},
	}
	app.Version = "0.6.0"
	app.EnableBashCompletion = true

	app.Flags = globalFlags
	app.Commands = []cli.Command{
		add,
		alertUnpaid,
		deleteUnpaid,
		disable,
		disableUnpaid,
		freeUser,
		generate,
		newYear,
		renable,
		renew,
		reset,
		resetShell,
		search,
		update,
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
