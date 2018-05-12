package rbldap

import (
	"errors"
	"time"

	"github.com/redbrick/rbldap/pkg/rbuser"
	"github.com/urfave/cli"
)

// Update a user in ldap
func Update(ctx *cli.Context) error {
	p := newPrompt()
	// Get User from arg if there prompt if not
	username := ""
	if ctx.NArg() > 0 {
		username = ctx.Args().First()
	} else {
		name, err := p.ReadString("Enter Username")
		if err != nil {
			return err
		}
		username = name
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

	user, err := rb.Search(filterAnd(filter("uid", username)))
	if user.UID == "" || err == nil {
		return errors.New("User not found")
	}
	// Prompt for details to change
	user.CN = p.Update("User's name", user.CN)
	user.Altmail = p.Update("User's email", user.Altmail)
	user.Course = p.Update("User's Course", user.Course)
	user.Year = p.UpdateInt("User's Year", user.Year)
	updatedBy, err := p.ReadUser("Updated By")
	if err != nil {
		return err
	}
	user.UpdatedBy = updatedBy
	user.LoginShell = p.UpdateShell(user.LoginShell)
	birthday, err := time.Parse("2006-01-02 15:04:00", p.Update("Users name", user.Birthday.Format("2006-01-02 15:04:00")))
	if err != nil {
		return err
	}
	user.Birthday = birthday
	return rb.Update(user)
}
