package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "userVhost"
	app.Usage = "apache conf generator for ldap"
	app.ArgsUsage = ""
	app.HideVersion = true
	app.EnableBashCompletion = true

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "user, u",
			Value: "cn=root,ou=ldap,o=redbrick",
			Usage: "ldap user, used for authentication",
		},

		cli.StringFlag{
			Name:  "host",
			Value: "ldap.internal",
			Usage: "ldap host to query",
		},

		cli.IntFlag{
			Name:  "port, p",
			Value: 389,
			Usage: "Port for ldap host",
		},

		cli.StringFlag{
			Name:     "password",
			Usage:    "password for the ldap server",
			FilePath: "/etc/ldap.secret",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:    "generate",
			Aliases: []string{"g"},
			Usage:   "generate list for uservhost macro",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "conf, c",
					Value: "./user_vhost_list.conf",
					Usage: "File to output conf too",
				},
			},
			Action: func(ctx *cli.Context) error {
				if ctx.NArg() != 0 {
					fmt.Fprintf(os.Stderr, "\n")
					return fmt.Errorf("Missing required arguments")
				}

				ldap := LdapConf{
					User:     ctx.String("user"),
					Password: ctx.String("password"),
					Host:     ctx.String("host"),
					Port:     ctx.Int("port"),
				}
				n, err := Generate(ldap, ctx.String("path"))
				if err != nil {
					return err
				}
				fmt.Printf("wrote %d bytes ./user_vhost_list.conf\n", n)
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
