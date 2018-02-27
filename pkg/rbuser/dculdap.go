package rbuser

import (
	"fmt"
	"regexp"
	"strconv"

	ldap "gopkg.in/ldap.v2"
)

// DcuLdap Server object used for connecting to DCU AD
type DcuLdap struct {
	*ldapConf
}

// NewDcuLdap create ldap connection to DCU AD
func NewDcuLdap(user, password, host string, port int) (*DcuLdap, error) {
	conf := &ldapConf{user: user, password: password, host: host, port: port}
	dcu := DcuLdap{conf}
	dcu.connect()
	return &dcu, nil
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
		course, year := splitCourseYear(entry.GetAttributeValue("l"))
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

func splitCourseYear(courseYear string) (string, int) {
	r, _ := regexp.Compile("([A-Z]+)")
	rYear, _ := regexp.Compile("([0-9]+)")

	year, _ := strconv.Atoi(rYear.FindString(courseYear))
	return r.FindString(courseYear), year
}
