package main

import (
	"github.com/redbrick/rb-ldap/internal/pkg/rbldap"
	"github.com/urfave/cli"
)

var search = cli.Command{
	Action:  rbldap.Search,
	Aliases: []string{"s"},
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "mail, altmail",
			Usage: "User email",
			Value: "*",
		},
		cli.StringFlag{
			Name:  "user, u, uid, nick, username",
			Usage: "User username",
			Value: "*",
		},
		cli.StringFlag{
			Name:  "id",
			Usage: "DCU id Number",
			Value: "*",
		},
		cli.StringFlag{
			Name:  "name, fullname",
			Usage: "User's fullname",
			Value: "*",
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
	Name:  "search",
	Usage: "Search ldap for user",
}
