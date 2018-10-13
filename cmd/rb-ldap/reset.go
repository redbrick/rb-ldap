package main

import (
	"github.com/redbrick/rb-ldap/internal/pkg/rbldap"
	"github.com/urfave/cli"
)

var reset = cli.Command{
	Action: rbldap.Reset,
	Name:   "reset",
	Usage:  "reset a users password",
}

var resetShell = cli.Command{
	Action: rbldap.ResetShell,
	Name:   "reset-shell",
	Usage:  "reset a users shell",
}
