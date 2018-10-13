package main

import (
	"github.com/redbrick/rb-ldap/internal/pkg/rbldap"
	"github.com/urfave/cli"
)

var freeUser = cli.Command{
	Action: rbldap.FreeUser,
	Name:   "free",
	Usage:  "Check if a username is free",
}
