package rbldap

import (
	"fmt"
	"os"
	"strings"

	"github.com/redbrick/rbldap/pkg/rbuser"
	"github.com/urfave/cli"
)

// Generate takes cli context and generates user vhost for rbuser
func Generate(ctx *cli.Context) error {
	rb, err := rbuser.NewRbLdap(
		ctx.String("user"),
		ctx.String("password"),
		ctx.String("host"),
		ctx.Int("port"),
	)
	if err != nil {
		return err
	}
	vhosts, err := rb.Generate()
	if err != nil {
		return err
	}
	if ctx.Bool("dry-run") {
		fmt.Print(strings.Join(vhosts, "\n"))
		return nil
	}
	file, err := os.Create(ctx.String("conf"))
	if err != nil {
		return err
	}
	defer file.Close()
	n, err := file.WriteString(strings.Join(vhosts, "\n"))
	if err != nil {
		return err
	}
	fmt.Printf("wrote %d bytes %s\n", n, ctx.String("conf"))
	return file.Sync()
}
