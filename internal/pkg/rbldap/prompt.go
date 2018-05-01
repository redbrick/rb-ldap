package rbldap

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
	"strconv"
	"strings"
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
	response, err := p.ReadString(message + " [" + user.Username + "]")
	if err != nil {
		return "", err
	}
	if response == "" {
		return user.Username, nil
	}
	return response, nil
}

func (p *prompt) confirm(message string) (bool, error) {
	res, err := p.ReadString(message + " [y/N]")
	if err != nil {
		return false, err
	}
	// Empty input (i.e. "\n")
	if len(res) < 2 {
		return false, nil
	}
	return strings.ToLower(res)[0] == 'y', nil
}
