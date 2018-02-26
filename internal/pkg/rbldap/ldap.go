package rbldap

import (
	"fmt"
	"os"

	"github.com/redbrick/rbldap/pkg/rbuser"
	"github.com/urfave/cli"
	ldap "gopkg.in/ldap.v2"
)

func newRBLdap(ctx *cli.Context) (*ldap.Conn, error) {
	if ctx.NArg() != 0 {
		fmt.Fprintf(os.Stderr, "\n")
		return &ldap.Conn{}, fmt.Errorf("Missing required arguments")
	}

	rb := rbuser.LdapConf{
		User:     ctx.String("user"),
		Password: ctx.String("password"),
		Host:     ctx.String("host"),
		Port:     ctx.Int("port"),
	}
	return rb.Connect()
}

func newDCULdap(ctx *cli.Context) (*ldap.Conn, error) {
	if ctx.NArg() != 0 {
		fmt.Fprintf(os.Stderr, "\n")
		return &ldap.Conn{}, fmt.Errorf("Missing required arguments")
	}

	dcu := rbuser.LdapConf{
		User:     ctx.String("dcu-user"),
		Password: ctx.String("dcu-password"),
		Host:     ctx.String("dcu-host"),
		Port:     ctx.Int("dcu-port"),
	}
	return dcu.Connect()
}
