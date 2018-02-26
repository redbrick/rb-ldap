package rbldap

import (
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli"
)

// Generate takes cli context and generates user vhost for rbuser
func Generate(ctx *cli.Context) error {
	l, err := newRbLdap(ctx)
	if err != nil {
		return err
	}
	vhosts, err := l.Generate()
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
