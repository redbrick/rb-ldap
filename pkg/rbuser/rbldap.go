package rbuser

import (
	"strconv"
	"strings"
	"time"

	ldap "gopkg.in/ldap.v2"
)

// RbLdap Server object used for connecting to server
type RbLdap struct {
	*LdapConf
}

// Search ldap for a given filter and return first user that matches
func (rb *RbLdap) Search(filter string) (RbUser, error) {
	sr, err := rb.Conn.Search(ldap.NewSearchRequest(
		"ou=accounts,o=redbrick",
		ldap.ScopeSingleLevel, ldap.NeverDerefAliases,
		0, 0, false, filter,
		[]string{"objectClass", "uid", "newbie", "cn", "altmail", "id", "course", "year",
			"yearsPaid", "updatedBy", "updated", "createdBy", "created", "birthday", "uidNumber",
			"gidNumber", "gecos", "loginShell", "homeDirectory", "userPassword", "host",
			"shadowLastChange"}, nil,
	))
	if err != nil {
		return RbUser{}, err
	}
	for _, entry := range sr.Entries {
		noob, _ := strconv.ParseBool(entry.GetAttributeValue("newbie"))
		dcuID, _ := strconv.Atoi(entry.GetAttributeValue("id"))
		year, _ := strconv.Atoi(entry.GetAttributeValue("year"))
		yearsPaid, _ := strconv.Atoi(entry.GetAttributeValue("yearsPaid"))
		uidNum, _ := strconv.Atoi(entry.GetAttributeValue("uidNumber"))
		gidNum, _ := strconv.Atoi(entry.GetAttributeValue("gidNumber"))
		updated, _ := time.Parse("2006-01-02 15:04:00", entry.GetAttributeValue("updated"))
		shadow, _ := time.Parse("2006-01-02 15:04:00", entry.GetAttributeValue("shadowLastChange"))
		created, _ := time.Parse("2006-01-02 15:04:00", entry.GetAttributeValue("created"))
		birthday, _ := time.Parse("2006-01-02 15:04:00", entry.GetAttributeValue("birthday"))
		return RbUser{
			UID:              entry.GetAttributeValue("uid"),
			ObjectClass:      entry.GetAttributeValue("objectClass"),
			Newbie:           noob,
			CN:               entry.GetAttributeValue("cn"),
			Altmail:          entry.GetAttributeValue("altmail"),
			ID:               dcuID,
			Course:           entry.GetAttributeValue("course"),
			Year:             year,
			YearsPaid:        yearsPaid,
			Updatedby:        entry.GetAttributeValue("updatedBy"),
			Updated:          updated,
			CreatedBy:        entry.GetAttributeValue("createdBy"),
			Created:          created,
			Birthday:         birthday,
			UIDNumber:        uidNum,
			GidNumber:        gidNum,
			Gecos:            entry.GetAttributeValue("gecos"),
			LoginShell:       entry.GetAttributeValue("loginShell"),
			HomeDirectory:    entry.GetAttributeValue("homeDirectory"),
			UserPassword:     entry.GetAttributeValue("userPassword"),
			Host:             strings.Split(entry.GetAttributeValue("host"), ","),
			ShadowLastChange: shadow,
		}, nil
	}
	return RbUser{}, err
}
