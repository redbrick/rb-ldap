package rbuser

import (
	"strconv"

	ldap "gopkg.in/ldap.v2"
)

// Generate user vhost conf from ldap
func Generate(l *ldap.Conn) ([]string, error) {
	var vhosts []string
	searchRequest := ldap.NewSearchRequest(
		"ou=accounts,o=redbrick",
		ldap.ScopeSingleLevel, ldap.NeverDerefAliases,
		0, 0, false, "(&)",
		[]string{"objectClass", "uid", "gidNumber"},
		nil,
	)
	sr, err := l.Search(searchRequest)
	if err != nil {
		return vhosts, err
	}
	for _, entry := range sr.Entries {
		group := entry.GetAttributeValue("objectClass")
		if group == "" {
			gidNum, conversionErr := strconv.Atoi(entry.GetAttributeValue("gidNumber"))
			if conversionErr != nil {
				return vhosts, conversionErr
			}
			group = gidToGroup(gidNum)
		}
		if group != "" && group != "redbrick" && group != "reserved" {
			u := RBUser{UID: entry.GetAttributeValue("uid"), ObjectClass: group}
			vhosts = append(vhosts, u.Vhost())
		}
	}
	return vhosts, nil
}

func gidToGroup(gid int) string {
	switch gid {
	case 100:
		return "committe"
	case 101:
		return "society"
	case 102:
		return "club"
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
