package main

import (
	"github.com/redbrick/rb-ldap/internal/pkg/rbldap"
	"github.com/urfave/cli"
)

var renew = cli.Command{
	Action:  rbldap.Renew,
	Aliases: []string{"r"},
	Name:    "renew",
	Usage:   "renew a LDAP user",
}
