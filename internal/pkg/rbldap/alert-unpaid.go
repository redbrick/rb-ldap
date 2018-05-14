package rbldap

import (
	"github.com/redbrick/rbldap/pkg/rbuser"
	"github.com/urfave/cli"
)

// AlertUnPaid emails all unpaid users warning about account diable
func AlertUnPaid(ctx *cli.Context) error {
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
	return rb.AlertUnPaid()
}
