package main

import "fmt"

type rbUser struct {
	Name    string
	Initial rune
	Group   string
}

func (u *rbUser) Vhost() string {
	return fmt.Sprintf("use VHost /storage/webtree/%s/%s %s %s %s", string(u.Initial), u.Name, u.Name, u.Group, u.Name)
}
