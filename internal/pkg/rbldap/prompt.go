package rbldap

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
	"strconv"
)

type prompt struct{}

func newPrompt() prompt {
	return prompt{}
}

func (p *prompt) ReadString(message string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(message + ": ")
	return reader.ReadString('\n')
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
