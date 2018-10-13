package main

import (
	"github.com/redbrick/rb-ldap/internal/pkg/rbldap"
	"github.com/urfave/cli"
)

var add = cli.Command{
	Action:  rbldap.Add,
	Aliases: []string{"a"},
	Name:    "add",
	Usage:   "Add user to ldap",
}
