package rbuser

import (
	"fmt"
	"regexp"
	"strconv"

	ldap "gopkg.in/ldap.v2"
)

// DcuLdap Server object used for connecting to dcu AD
type DcuLdap struct {
	*LdapConf
}

// Search dcu ldap for a given filter and return first user that matches
func (dcu *DcuLdap) Search(filter string) (RbUser, error) {
	sr, err := dcu.Conn.Search(ldap.NewSearchRequest(
		"o=ad,o=dcu,o=ie",
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases,
		0, 0, false, filter,
		[]string{"employeeNumber", "givenName", "sn", "gecos", "mail", "l"}, nil,
	))
	if err != nil {
		return RbUser{}, err
	}
	for _, entry := range sr.Entries {
		dcuID, _ := strconv.Atoi(entry.GetAttributeValue("employeeNumber"))
		course, year := courseYear(entry.GetAttributeValue("l"))
		return RbUser{
			CN:      fmt.Sprintf("%s %s", entry.GetAttributeValue("givenname"), entry.GetAttributeValue("sn")),
			Altmail: entry.GetAttributeValue("mail"),
			ID:      dcuID,
			Course:  course,
			Year:    year,
		}, nil
	}
	return RbUser{}, err
}

func courseYear(whyNotTheSame string) (string, int) {
	r, _ := regexp.Compile("([A-Z]+)")
	rYear, _ := regexp.Compile("([0-9]+)")

	year, _ := strconv.Atoi(rYear.FindString(whyNotTheSame))
	return r.FindString(whyNotTheSame), year
}
