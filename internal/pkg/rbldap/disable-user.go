package rbldap

import (
	"github.com/redbrick/rb-ldap/pkg/rbuser"
	"github.com/urfave/cli"
)

// DisableUser disable a users ldap account
func DisableUser(ctx *cli.Context) error {
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
	if ctx.GlobalBool("dry-run") {
		return errNotImplemented
	}
	return rb.DisableUser(user)
}

// RenableUser renable a users ldap account
func RenableUser(ctx *cli.Context) error {
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
	if ctx.GlobalBool("dry-run") {
		return errNotImplemented
	}
	return rb.RenableUser(user)
}
