package main

import (
	"github.com/redbrick/rb-ldap/internal/pkg/rbldap"
	"github.com/urfave/cli"
)

var generate = cli.Command{
	Action:  rbldap.Generate,
	Aliases: []string{"g"},
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "conf, c",
			Usage: "Output configuration `FILE`",
			Value: "./user_vhost_list.conf",
		},
	},
	Name:  "generate",
	Usage: "generate list for uservhost macro",
}
