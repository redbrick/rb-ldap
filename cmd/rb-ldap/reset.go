package main

import (
	"github.com/redbrick/rbldap/internal/pkg/rbldap"
	"github.com/urfave/cli"
)

var reset = cli.Command{
	Action: rbldap.Reset,
	Name:   "reset",
	Usage:  "reset a users password",
}
