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
	vhosts, err := rb.Generate()
	if err != nil {
		return err
	}
	if ctx.GlobalBool("dry-run") {
		fmt.Print(strings.Join(vhosts, "\n"))
		return nil
	}
	file, err := os.Create(ctx.String("conf"))
	if err != nil {
		return err
	}
	defer file.Close()
	n, err := file.WriteString("# WARNING, This File is auto Generated do NOT edit.\n# Please go read the docs. https://docs.redbrick.dcu.ie/web/apache\n" + strings.Join(vhosts, "\n"))
	if err != nil {
		return err
	}
	fmt.Printf("wrote %d bytes %s\n", n, ctx.String("conf"))
	return file.Sync()
}
