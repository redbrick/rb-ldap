package rbldap

import (
	"errors"
	"time"

	"github.com/redbrick/rbldap/pkg/rbuser"
	"github.com/urfave/cli"
)

// Update a user in ldap
func Update(ctx *cli.Context) error {
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
		return errors.New("User not found")
	}
	// Prompt for details to change
	p := newPrompt()
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
	birthday, err := time.Parse("2006-01-02 15:04:05", p.Update("Users name", user.Birthday.Format("2006-01-02 15:04:05")))
	if err != nil {
		return err
	}
	user.Birthday = birthday
	if ctx.GlobalBool("dry-run") {
		return user.PrettyPrint()
	}
	return rb.Update(user)
}

// Get Username from arg if there and prompt for it if not
func getUsername(args cli.Args) (string, error) {
	if len(args) > 0 {
		return args.First(), nil
	}
	p := newPrompt()
	return p.ReadString("Enter Username")
}
