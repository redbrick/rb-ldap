package main

import (
	"github.com/redbrick/rbldap/internal/pkg/rbldap"
	"github.com/urfave/cli"
)

var newYear = cli.Command{
	Action:      rbldap.NewYear,
	Category:    "Batch Commands",
	Name:        "new-year",
	Usage:       "Decriment Years Paid of all users to 1",
	Description: "Migrate all users to no longer be marked as newbies and mark all users as unpaided. To Be run at the beginning of each year prior to C&S",
}
