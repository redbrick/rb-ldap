package main

import (
	"github.com/redbrick/rb-ldap/internal/pkg/rbldap"
	"github.com/urfave/cli"
)

var disable = cli.Command{
	Action:  rbldap.DisableUser,
	Name:    "disable",
	Aliases: []string{"disuser"},
	Usage:   "Disable a Users ldap account",
}

var renable = cli.Command{
	Action:  rbldap.RenableUser,
	Name:    "renable",
	Aliases: []string{"reuser"},
	Usage:   "Renable a Users ldap account",
}
