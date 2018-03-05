package rbldap

import (
	"bytes"
	"fmt"
	"os"
	"regexp"

	"github.com/redbrick/rbldap/pkg/rbuser"
	"github.com/urfave/cli"
)

// Search command for cli app
func Search(ctx *cli.Context) error {
	if ctx.NArg() != 0 {
		fmt.Fprintf(os.Stderr, "\n")
		return fmt.Errorf("Missing required arguments")
	}
	mail := filter("altmail", ctx.String("mail"))
	id := filter("id", ctx.String("id"))
	re := regexp.MustCompile(`\ `)
	name := re.ReplaceAllString(ctx.String("name"), `*$1*$2*`)
	if ctx.Bool("dcu") {
		dcu, err := rbuser.NewDcuLdap(
			ctx.String("dcu-user"),
			ctx.String("dcu-password"),
			ctx.String("dcu-host"),
			ctx.Int("dcu-port"),
		)
		if err != nil {
			return err
		}
		user, searchErr := dcu.Search(filterAnd(
			filter("displayName", name),
			filter("cn", ctx.String("user")),
			id),
		)
		if searchErr != nil {
			return searchErr
		}
		return user.PrettyPrint()
	}
	rb, err := rbuser.NewRbLdap(
		ctx.String("user"),
		ctx.String("password"),
		ctx.String("host"),
		ctx.Int("port"),
	)
	if err != nil {
		return err
	}
	noob := ""
	if ctx.Bool("noob") {
		noob = "(newbie=TRUE)"
	}
	user, err := rb.Search(filterAnd(
		filter("cn", name),
		filterOr(
			filter("uid", ctx.String("user")),
			filter("gecos", ctx.String("user")),
		), id, mail, noob),
	)
	if err != nil {
		return err
	}
	return user.PrettyPrint()
}

func filter(key, search string) string {
	return fmt.Sprintf("(%s=%s)", key, search)
}

func filterAnd(args ...string) string {
	return filterJoin("&", args)
}

func filterOr(args ...string) string {
	return filterJoin("|", args)
}

func filterJoin(join string, args []string) string {
	var buffer bytes.Buffer
	buffer.WriteString("(")
	buffer.WriteString(join)
	for _, filter := range args {
		buffer.WriteString(filter)
	}
	buffer.WriteString(")")
	return buffer.String()
}
