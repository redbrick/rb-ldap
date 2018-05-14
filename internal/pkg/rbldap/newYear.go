package rbldap

import (
	"github.com/redbrick/rbldap/pkg/rbuser"
	"github.com/urfave/cli"
)

// NewYear run new year migration on ldap
func NewYear(ctx *cli.Context) error {
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
	p := newPrompt()
	username, err := p.ReadUser("Created by")
	if err != nil {
		return err
	}
	return rb.NewYear(username)
}
