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

var deleteUnpaid = cli.Command{
	Action:   rbldap.DeleteUnPaid,
	Category: "Batch Commands",
	Name:     "delete-unpaid",
	Usage:    "Delete all unpaid users accounts that are outside their grace period",
}

var disableUnpaid = cli.Command{
	Action:   rbldap.DisableUnPaid,
	Category: "Batch Commands",
	Name:     "disable-unpaid",
	Usage:    "Diable all unpaid users accounts",
}
