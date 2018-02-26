package rbldap

import (
	"fmt"
	"os"

	"github.com/redbrick/rbldap/pkg/rbuser"
	"github.com/urfave/cli"
)

func newRbLdap(ctx *cli.Context) (*rbuser.RbLdap, error) {
	if ctx.NArg() != 0 {
		fmt.Fprintf(os.Stderr, "\n")
		return &rbuser.RbLdap{}, fmt.Errorf("Missing required arguments")
	}

	rb := rbuser.RbLdap{
		User:     ctx.String("user"),
		Password: ctx.String("password"),
		Host:     ctx.String("host"),
		Port:     ctx.Int("port"),
	}
	rb.Connect()
	return &rb, nil
}

func newDcuLdap(ctx *cli.Context) (*rbuser.RbLdap, error) {
	if ctx.NArg() != 0 {
		fmt.Fprintf(os.Stderr, "\n")
		return &rbuser.RbLdap{}, fmt.Errorf("Missing required arguments")
	}

	dcu := rbuser.RbLdap{
		User:     ctx.String("dcu-user"),
		Password: ctx.String("dcu-password"),
		Host:     ctx.String("dcu-host"),
		Port:     ctx.Int("dcu-port"),
	}
	dcu.Connect()
	return &dcu, nil
}
