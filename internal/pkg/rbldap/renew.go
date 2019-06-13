package rbldap

import (
	"fmt"

	"github.com/redbrick/rb-ldap/pkg/rbuser"
	"github.com/urfave/cli"
)

// Renew a users LDAP account
func Renew(ctx *cli.Context) error {
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
	} else if user.YearsPaid == 1 {
		return nil
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
	dcuUser, err := dcu.Search(filterAnd(filter("employeenumber", string(user.ID))))
	if err != nil {
		return err
	}
	user.Year = dcuUser.Year
	user.Course = dcuUser.Course
	user.UserType = dcuUser.UserType
	p := newPrompt()
	updatedBy, err := p.ReadUser("Updated by")
	if err != nil {
		return err
	}
	user.UpdatedBy = updatedBy
	newHome := fmt.Sprintf("/home/%s/%s/%s", user.UserType, []rune(user.UID)[0], user.UID)
	if ctx.GlobalBool("dry-run") {
		user.HomeDirectory = newHome
		return user.PrettyPrint()
	}
	if err := user.MigrateHome(newHome); err != nil {
		return err
	}
	return rb.Renew(user)
}
