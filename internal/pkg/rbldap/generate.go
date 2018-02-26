package rbldap

import (
	"fmt"

	"github.com/redbrick/rbldap/pkg/rbuser"
	"github.com/urfave/cli"
)

// Generate takes cli context and generates user vhost for rbuser
func Generate(ctx *cli.Context) error {
	l, err := newRBLdap(ctx)
	if err != nil {
		return err
	}
	path := ctx.String("conf")
	n, genErr := rbuser.Generate(l, path)
	fmt.Printf("wrote %d bytes %s\n", n, path)
	return genErr
}
