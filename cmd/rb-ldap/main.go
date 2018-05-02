package main

import (
	"fmt"
	"os"

	"github.com/redbrick/rbldap/internal/pkg/rbldap"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "rb-ldap"
	app.Usage = "Command line interface for Redbrick LDAP"
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Cian Butler",
			Email: "butlerx@redbrick.dcu.ie",
		},
	}
	app.Version = "0.2.0"
	app.EnableBashCompletion = true

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "user, u",
			Value: "cn=root,ou=ldap,o=redbrick",
			Usage: "ldap user, used for authentication",
		},

		cli.StringFlag{
			Name:  "dcu-user",
			Value: "CN=rblookup,OU=Service Accounts,DC=ad,DC=dcu,DC=ie",
			Usage: "Active Directory user for DCU, used for authentication",
		},

		cli.StringFlag{
			Name:  "host",
			Value: "ldap.internal",
			Usage: "ldap host to query",
		},

		cli.StringFlag{
			Name:  "dcu-host",
			Value: "ad.dcu.ie",
			Usage: "DCU Active Directory host to query",
		},

		cli.IntFlag{
			Name:  "port, p",
			Value: 389,
			Usage: "Port for ldap host",
		},

		cli.IntFlag{
			Name:  "dcu-port",
			Value: 389,
			Usage: "Port for DCU Active Directory host",
		},

		cli.StringFlag{
			Name:     "password",
			Usage:    "password for the ldap server",
			FilePath: "/etc/ldap.secret",
		},

		cli.StringFlag{
			Name:     "dcu-password",
			Usage:    "password for the DCU ldap server",
			FilePath: "/etc/dcu_ldap.secret",
		},

		cli.StringFlag{
			Name:     "smtp",
			Usage:    "smtp server to send email with",
			FilePath: "mailhost.redbrick.dcu.ie",
		},

		cli.BoolFlag{
			Name:  "dry-run",
			Usage: "output to console rather then file",
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
					Usage: "Output configuration `FILE`",
				},
			},
			Action: rbldap.Generate,
		},
		{
			Name:    "search",
			Aliases: []string{"s"},
			Usage:   "Search ldap for user",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "mail, altmail",
					Value: "*",
					Usage: "User email",
				},
				cli.StringFlag{
					Name:  "user, u, uid, nick, username",
					Value: "*",
					Usage: "User username",
				},
				cli.StringFlag{
					Name:  "id",
					Value: "*",
					Usage: "DCU id Number",
				},
				cli.StringFlag{
					Name:  "name, fullname",
					Value: "*",
					Usage: "User's fullname",
				},
				cli.BoolFlag{
					Name:  "newbie, noob",
					Usage: "filter for new users",
				},
				cli.BoolFlag{
					Name:  "dcu, DCU",
					Usage: "Query DCU Active Directory",
				},
			},
			Action: rbldap.Search,
		},
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "Add user to ldap",
			Action:  rbldap.Add,
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
