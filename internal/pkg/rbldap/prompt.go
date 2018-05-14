package rbldap

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"strconv"
	"strings"

	"github.com/urfave/cli"
)

type prompt struct {
	reader *bufio.Reader
}

func newPrompt() prompt {
	return prompt{bufio.NewReader(os.Stdin)}
}

func (p *prompt) ReadString(message string) (string, error) {
	fmt.Print(message + ": ")
	res, err := p.reader.ReadString('\n')
	return strings.TrimSpace(res), err
}

func (p *prompt) ReadInt(message string) (int, error) {
	msg, err := p.ReadString(message)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(msg)
}

func (p *prompt) ReadUser(message string) (string, error) {
	user, err := user.Current()
	if err != nil {
		return "", err
	}
	res, err := p.ReadString(message + " [" + user.Username + "]")
	if err != nil || len(res) < 2 {
		return user.Username, err
	}
	return res, nil
}

func (p *prompt) Confirm(message string) (bool, error) {
	res, err := p.ReadString(message + " [y/N]")
	if err != nil || len(res) < 2 {
		return false, err
	}
	return strings.ToLower(res)[0] == 'y', nil
}

func (p *prompt) Update(message, original string) string {
	res, err := p.ReadString(message + " [" + original + "]")
	if err != nil || len(res) < 2 {
		return original
	}
	return res
}

func (p *prompt) UpdateInt(message string, original int) int {
	res := p.Update(message, strconv.Itoa(original))
	result, err := strconv.Atoi(res)
	if err != nil {
		return original
	}
	return result
}

func (p *prompt) UpdateShell(original string) string {
	content, err := ioutil.ReadFile("/etc/shells")
	if err != nil {
		return original
	}
	var shells []string
	for _, str := range strings.Split(string(content), "\n") {
		// Check for empty lines or comments
		if str != "" && string([]rune(str)[0]) != "#" {
			shells = append(shells, str)
		}
	}
	for {
		res, err := p.ReadString("User's Shell" + " [" + original + "]")
		if err != nil || !validShell(shells, res) {
			fmt.Printf("Your Shell choice is not valid.\n Available shells are: %+q\n", shells)
		} else if len(res) < 2 {
			return original
		}
		return res
	}
}

func validShell(shells []string, check string) bool {
	for _, shell := range shells {
		if shell == check {
			return true
		}
	}
	return false
}

// Get Username from arg if there and prompt for it if not
func getUsername(args cli.Args) (string, error) {
	if len(args) > 0 {
		return args.First(), nil
	}
	p := newPrompt()
	return p.ReadString("Enter Username")
}
