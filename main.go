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

type rbUser struct {
	Name    string
	Initial rune
	Group   string
}

func (u *rbUser) Vhost() string {
	return fmt.Sprintf("use VHost /storage/webtree/%s/%s %s %s %s", string(u.Initial), u.Name, u.Name, u.Group, u.Name)
}

func main() {
	binduser := flag.String("user", "cn=root,ou=ldap,o=redbrick", "ldap user")
	host := flag.String("host", "ldap.internal", "ldap host")
	port := flag.Int("port", 389, "ldap port")
	path := flag.String("path", "./user_vhost_list.conf", "file to output too")
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		log.Fatal("password secret must be specified")
	}

	secret, err := ioutil.ReadFile(args[0])
	if err != nil {
		log.Fatal(err)
	}
	bindpassword := strings.Join(strings.Fields(string(secret)), " ")
	n, err := Generate(*binduser, bindpassword, *host, *port, *path)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("wrote %d bytes ./user_vhost_list.conf\n", n)
}

// Generate user vhost conf from ldap
func Generate(user string, pass string, host string, port int, output string) (int, error) {
	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		return 0, err
	}
	defer l.Close()

	err = l.Bind(user, pass)
	if err != nil {
		return 0, err
	}

	searchRequest := ldap.NewSearchRequest(
		"ou=accounts,o=redbrick",
		ldap.ScopeSingleLevel, ldap.NeverDerefAliases,
		0, 0, false, "(&)",
		[]string{"objectClass", "uid", "gidNumber"},
		nil,
	)
	sr, err := l.Search(searchRequest)
	if err != nil {
		return 0, err
	}
	var vhosts []string
	for _, entry := range sr.Entries {
		uid := entry.GetAttributeValue("uid")
		group := entry.GetAttributeValue("objectClass")
		if group == "" {
			group = gidToGroup(entry)
		}
		if group != "" && group != "redbrick" && group != "reserved" {
			u := rbUser{uid, []rune(uid)[0], group}
			vhosts = append(vhosts, u.Vhost())
		}
	}
	f, err := os.Create(output)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	n, err := f.WriteString(strings.Join(vhosts, "\n"))
	f.Sync()
	return n, err
}

func gidToGroup(entry *ldap.Entry) string {
	gid := entry.GetAttributeValue("gidNumber")
	gidNum, err := strconv.Atoi(gid)
	if err != nil {
		return ""
	}
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
