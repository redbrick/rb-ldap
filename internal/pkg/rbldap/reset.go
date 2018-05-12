package rbldap

import (
	"github.com/redbrick/rbldap/pkg/rbuser"
	"github.com/urfave/cli"
)

// Reset a users LDAP Password
func Reset(ctx *cli.Context) error {
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
	return rb.ResetPasswd(username)
}
