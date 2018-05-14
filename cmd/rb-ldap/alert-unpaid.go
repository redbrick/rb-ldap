package main

import (
	"github.com/redbrick/rbldap/internal/pkg/rbldap"
	"github.com/urfave/cli"
)

var alertUnpaid = cli.Command{
	Action:   rbldap.AlertUnPaid,
	Category: "Batch Commands",
	Name:     "alert-unpaid",
	Usage:    "Alert all unpaid users that their accounts will be disabled",
}
