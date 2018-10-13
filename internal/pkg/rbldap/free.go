package rbldap

import (
	"fmt"

	"github.com/redbrick/rb-ldap/pkg/rbuser"
	"github.com/urfave/cli"
)

// FreeUser Check if a user name is free
func FreeUser(ctx *cli.Context) error {
	rb, err := rbuser.NewRbLdap(
		ctx.GlobalString("user"),
		ctx.GlobalString("password"),
		ctx.GlobalString("host"),
		ctx.GlobalInt("port"),
		ctx.GlobalString("smtp"),
	)
	if err != nil {
		return err
	}
	defer rb.Conn.Close()
	username, err := getUsername(ctx.Args())
	if err != nil {
		return err
	}
	user, err := rb.SearchUser(filterAnd(filter("uid", username)))
	if user.UID == "" || err != nil {
		fmt.Printf("%s is free\n", username)
		return nil
	}
	return errDuplicateUser
}
