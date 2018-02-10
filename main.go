package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	ldap "gopkg.in/ldap.v2"
)

type user struct {
	Name    string
	Initial rune
	Group   string
}

func (u *user) Vhost() string {
	return fmt.Sprintf("use VHost /storage/webtree/%s/%s %s %s %s", string(u.Initial), u.Name, u.Name, u.Group, u.Name)
}

func main() {
	binduser := flag.String("user", "cn=root,ou=ldap,o=redbrick", "ldap user")
	host := flag.String("host", "ldap.internal", "ldap host")
	port := flag.Int("port", 389, "ldap port")
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		log.Fatal("password secret must be specified")
	}

	secret, err := ioutil.ReadFile(args[0])
	check(err)
	bindpassword := strings.Join(strings.Fields(string(secret)), " ")

	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", *host, *port))
	check(err)
	defer l.Close()

	err = l.Bind(*binduser, bindpassword)
	check(err)

	searchRequest := ldap.NewSearchRequest(
		"ou=accounts,o=redbrick",
		ldap.ScopeSingleLevel, ldap.NeverDerefAliases,
		0, 0, false, "(&)",
		[]string{"objectClass", "uid", "gidNumber"},
		nil,
	)
	sr, err := l.Search(searchRequest)
	check(err)
	var vhosts []string
	for _, entry := range sr.Entries {
		uid := entry.GetAttributeValue("uid")
		group := entry.GetAttributeValue("objectClass")
		if group == "" {
			group = gidToGroup(entry)
		}
		if group != "" && group != "redbrick" && group != "reserved" {
			u := user{uid, []rune(uid)[0], group}
			vhosts = append(vhosts, u.Vhost())
		}
	}
	f, err := os.Create("./user_vhost_list.conf")
	check(err)
	defer f.Close()
	n, err := f.WriteString(strings.Join(vhosts, "\n"))
	fmt.Printf("wrote %d bytes ./user_vhost_list.conf\n", n)
	f.Sync()
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func gidToGroup(entry *ldap.Entry) string {
	gid := entry.GetAttributeValue("gidNumber")
	gidNum, err := strconv.Atoi(gid)
	check(err)
	switch gidNum {
	case 100:
		return "committe"
	case 101:
		return "society"
	case 102:
		return "club"
	case 103:
		return "member"
	case 105:
		return "founder"
	case 107:
		return "associat"
	case 109:
		return "staff"
	case 1016:
		return "intersoc"
	case 1017:
		return "redbrick"
	case 1014:
		return "projects"
	case 31382:
		return "dcu"
	default:
		return "member"
	}
}
