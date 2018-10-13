package main

import (
	"github.com/redbrick/rb-ldap/internal/pkg/rbldap"
	"github.com/urfave/cli"
)

var update = cli.Command{
	Action:  rbldap.Update,
	Aliases: []string{"u", "edit"},
	Name:    "update",
	Usage:   "Update a user in ldap",
}
