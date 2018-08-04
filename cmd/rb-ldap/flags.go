package main

import "github.com/urfave/cli"

var globalFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "user, u",
		Usage: "ldap user, used for authentication",
		Value: "cn=root,ou=ldap,o=redbrick",
	},

	cli.StringFlag{
		Name:  "dcu-user",
		Usage: "Active Directory user for DCU, used for authentication",
		Value: "CN=rblookup,OU=Service Accounts,DC=ad,DC=dcu,DC=ie",
	},

	cli.StringFlag{
		Name:  "host",
		Usage: "ldap host to query",
		Value: "ldap.internal",
	},

	cli.StringFlag{
		Name:  "dcu-host",
		Usage: "DCU Active Directory host to query",
		Value: "ad.dcu.ie",
	},

	cli.IntFlag{
		Name:  "port, p",
		Usage: "Port for ldap host",
		Value: 389,
	},

	cli.IntFlag{
		Name:  "dcu-port",
		Usage: "Port for DCU Active Directory host",
		Value: 389,
	},

	cli.StringFlag{
		FilePath: "/etc/ldap.secret",
		Name:     "password",
		Usage:    "password for the ldap server",
	},

	cli.StringFlag{
		FilePath: "/etc/dcu_ldap.secret",
		Name:     "dcu-password",
		Usage:    "password for the DCU ldap server",
	},

	cli.StringFlag{
		Value: "mailhost.redbrick.dcu.ie",
		Name:  "smtp",
		Usage: "smtp server to send email with",
	},

	cli.BoolFlag{
		Name:  "dry-run",
		Usage: "output to console rather then file",
	},
}
