package rbldap

import (
	"github.com/redbrick/rbldap/pkg/rbuser"
	"github.com/urfave/cli"
)

// AlertUnPaid emails all unpaid users warning about account diable
func AlertUnPaid(ctx *cli.Context) error {
	if ctx.GlobalBool("dry-run") {
		return errNotImplemented
	}
	p := newPrompt()
	if confirm, err := p.Confirm("Email All Unpaid Users"); !confirm || err != nil {
		return err
	}
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

// DisableUnPaid diable all unpaid users
func DisableUnPaid(ctx *cli.Context) error {
	if ctx.GlobalBool("dry-run") {
		return errNotImplemented
	}
	p := newPrompt()
	if confirm, err := p.Confirm("Disable All Unpaid Users"); !confirm || err != nil {
		return err
	}
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
	admin, err := p.ReadUser("Disabled by")
	if err != nil {
		return err
	}
	return rb.DisableUnPaid(admin)
}

// DeleteUnPaid emails all unpaid users warning about account diable
func DeleteUnPaid(ctx *cli.Context) error {
	if ctx.GlobalBool("dry-run") {
		return errNotImplemented
	}
	p := newPrompt()
	if confirm, err := p.Confirm("Delete All Unpaid Users, THIS CANNOT BE UNDONE"); !confirm || err != nil {
		return err
	}
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
	return rb.DeleteUnPaid()
}
