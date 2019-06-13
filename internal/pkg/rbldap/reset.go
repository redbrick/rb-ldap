package rbldap

import (
	"github.com/redbrick/rb-ldap/pkg/rbuser"
	"github.com/urfave/cli"
)

// Reset a users LDAP Password
func Reset(ctx *cli.Context) error {
	if ctx.GlobalBool("dry-run") {
		return errNotImplemented
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
	username, err := getUsername(ctx.Args())
	if err != nil {
		return err
	}
	return rb.ResetPasswd(username)
}

// ResetShell a users shell
func ResetShell(ctx *cli.Context) error {
	if ctx.GlobalBool("dry-run") {
		return errNotImplemented
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
	username, err := getUsername(ctx.Args())
	if err != nil {
		return err
	}
	user, err := rb.SearchUser(filterAnd(filter("uid", username)))
	if user.UID == "" || err == nil {
		return errUser404
	}
	p := newPrompt()
	updatedBy, err := p.ReadUser("Updated by")
	if err != nil {
		return err
	}
	user.UpdatedBy = updatedBy
	return rb.ResetShell(user)
}
