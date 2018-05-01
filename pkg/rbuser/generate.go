package rbuser

import (
	"strconv"

	ldap "gopkg.in/ldap.v2"
)

// Generate user vhost conf from ldap
func (l *RbLdap) Generate() ([]string, error) {
	var vhosts []string
	searchRequest := ldap.NewSearchRequest(
		"ou=accounts,o=redbrick",
		ldap.ScopeSingleLevel, ldap.NeverDerefAliases,
		0, 0, false, "(&)",
		[]string{"objectClass", "uid", "gidNumber"},
		nil,
	)
	sr, err := l.Conn.Search(searchRequest)
	if err != nil {
		return vhosts, err
	}
	for _, entry := range sr.Entries {
		group := entry.GetAttributeValue("objectClass")
		if group == "" {
			gidNum, err := strconv.Atoi(entry.GetAttributeValue("gidNumber"))
			if err != nil {
				return vhosts, err
			}
			group = gidToGroup(gidNum)
		}
		if group != "" && group != "redbrick" && group != "reserved" {
			u := RbUser{UID: entry.GetAttributeValue("uid"), ObjectClass: group}
			vhosts = append(vhosts, u.Vhost())
		}
	}
	return vhosts, nil
}
