package rbldap

import (
	"errors"

	"github.com/redbrick/rbldap/pkg/rbuser"
	"github.com/urfave/cli"
)

// Add a user to ldap
func Add(ctx *cli.Context) error {
	p := newPrompt()
	// check username free
	username, err := p.ReadString("Enter Username")
	if err != nil {
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
	foundUser, err := rb.Search(filterAnd(filter("uid", username)))
	if foundUser.UID != "" || err != nil {
		return errors.New("User Already exists")
	}
	// search dcu for id number and create RbUser
	id, err := p.ReadInt("Enter Student ID Number")
	if err != nil {
		return err
	}
	dcu, err := rbuser.NewDcuLdap(
		ctx.GlobalString("dcu-user"),
		ctx.GlobalString("dcu-password"),
		ctx.GlobalString("dcu-host"),
		ctx.GlobalInt("dcu-port"),
	)
	if err != nil {
		return err
	}
	defer dcu.Conn.Close()
	newUser, err := dcu.Search(filterAnd(filter("employeenumber", string(id))))
	if err != nil {
		return err
	}
	newUser.UID = username
	newUser.UserType = "member"
	createdBy, err := p.ReadUser("Created by")
	if err != nil {
		return err
	}
	newUser.CreatedBy = createdBy
	mailUser, err := p.confirm("Mail user Login info")
	if err != nil {
		return err
	}
	if ctx.GlobalBool("dry-run") {
		return newUser.PrettyPrint()
	}
	return rb.Add(newUser, mailUser)
}
