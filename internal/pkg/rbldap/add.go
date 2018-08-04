package rbldap

import (
	"strconv"

	"github.com/redbrick/rbldap/pkg/rbuser"
	"github.com/urfave/cli"
)

// Add a user to ldap
func Add(ctx *cli.Context) error {
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
	foundUser, err := rb.SearchUser(filterAnd(filter("uid", username)))
	if foundUser.UID != "" || err != nil {
		return errDuplicateUser
	}
	// search dcu for id number and create RbUser
	p := newPrompt()
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
	newUser, err := dcu.Search(filterAnd(filter("employeenumber", strconv.Itoa(id))))
	if err != nil {
		return err
	}
	newUser.UID = username
	createdBy, err := p.ReadUser("Created by")
	if err != nil {
		return err
	}
	newUser.CreatedBy = createdBy
	newUser.Host = []string{"azazel", "pygmalion"}
	mailUser, err := p.Confirm("Mail user Login info")
	if err != nil {
		return err
	}
	if ctx.GlobalBool("dry-run") {
		return newUser.PrettyPrint()
	}
	return rb.Add(newUser, mailUser)
}
