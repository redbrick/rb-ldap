package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	ldap "gopkg.in/ldap.v2"
)

// Generate user vhost conf from ldap
func Generate(conf LdapConf, output string) (int, error) {
	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", conf.Host, conf.Port))
	if err != nil {
		return 0, err
	}
	defer l.Close()

	err = l.Bind(conf.User, conf.Password)
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
			gidNum, conversionErr := strconv.Atoi(entry.GetAttributeValue("gidNumber"))
			if conversionErr != nil {
				return 0, conversionErr
			}
			group = gidToGroup(gidNum)
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
	if err != nil {
		return 0, err
	}
	n, err := f.WriteString(strings.Join(vhosts, "\n"))
	if err != nil {
		return n, err
	}
	err = f.Sync()
	return n, err
}

func gidToGroup(gid int) string {
	switch gid {
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
