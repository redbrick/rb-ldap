package rbuser

import (
	"errors"
	"regexp"
	"strconv"
	"strings"

	ldap "gopkg.in/ldap.v2"
)

// DcuLdap Server object used for connecting to DCU AD
type DcuLdap struct {
	*ldapConf
}

// NewDcuLdap create ldap connection to DCU AD
func NewDcuLdap(user, password, host string, port int) (*DcuLdap, error) {
	dcu := DcuLdap{
		&ldapConf{
			user:     user,
			password: password,
			host:     host,
			port:     port,
		},
	}
	return &dcu, dcu.connect()
}

// Search dcu ldap for a given filter and return first user that matches
func (dcu *DcuLdap) Search(filter string) (RbUser, error) {
	sr, err := dcu.Conn.Search(ldap.NewSearchRequest(
		"o=ad,o=dcu,o=ie",
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases,
		0, 0, false, filter,
		[]string{"employeeNumber", "displayName", "mail", "physicalDeliveryOfficeName", "distinguishedName"}, nil,
	))
	if err != nil {
		return RbUser{}, err
	}
	for _, entry := range sr.Entries {
		dcuID, _ := strconv.Atoi(entry.GetAttributeValue("employeeNumber"))
		course, year := splitCourseYear(entry.GetAttributeValue("physicalDeliveryOfficeName"))
		userType, userTypeErr := getUserType(entry.GetAttributeValue("distinguishedName"))
		if userTypeErr != nil {
			return RbUser{}, userTypeErr
		}
		return RbUser{
			CN:       entry.GetAttributeValue("displayName"),
			Altmail:  entry.GetAttributeValue("mail"),
			UserType: userType,
			ID:       dcuID,
			Course:   course,
			Year:     year,
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

func getUserType(dn string) (string, error) {
	temp := strings.Split(dn, ",")
	switch userType := strings.Split(temp[len(temp)-4], "=")[1]; userType {
	case "Student":
		return "member", nil
	case "Staff":
		return "staff", nil
	case "Alumni":
		return "associat", nil
	default:
		return "", errors.New("Unknown UserType")
	}
}
