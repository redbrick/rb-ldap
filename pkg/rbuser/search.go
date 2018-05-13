package rbuser

import (
	"strconv"
	"time"

	ldap "gopkg.in/ldap.v2"
)

const timeLayout = "2006-01-02 15:04:05"

// SearchUser ldap for a given filter and return first user that matches
func (rb *RbLdap) SearchUser(filter string) (RbUser, error) {
	users, err := rb.SearchUsers(filter)
	return users[0], err
}

// SearchUsers ldap for a given filter and return all users that matches
func (rb *RbLdap) SearchUsers(filter string) ([]RbUser, error) {
	var users []RbUser
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
		return users, err
	}
	for _, entry := range sr.Entries {
		noob, _ := strconv.ParseBool(entry.GetAttributeValue("newbie"))
		dcuID, _ := strconv.Atoi(entry.GetAttributeValue("id"))
		year, _ := strconv.Atoi(entry.GetAttributeValue("year"))
		yearsPaid, _ := strconv.Atoi(entry.GetAttributeValue("yearsPaid"))
		uidNum, _ := strconv.Atoi(entry.GetAttributeValue("uidNumber"))
		gidNum, _ := strconv.Atoi(entry.GetAttributeValue("gidNumber"))
		updated, _ := time.Parse(timeLayout, entry.GetAttributeValue("updated"))
		shadow, _ := strconv.Atoi(entry.GetAttributeValue("shadowLastChange"))
		created, _ := time.Parse(timeLayout, entry.GetAttributeValue("created"))
		birthday, _ := time.Parse(timeLayout, entry.GetAttributeValue("birthday"))
		users = append(users, RbUser{
			UID:              entry.GetAttributeValue("uid"),
			UserType:         entry.GetAttributeValue("objectClass"),
			ObjectClass:      entry.GetAttributeValues("objectClass"),
			Newbie:           noob,
			CN:               entry.GetAttributeValue("cn"),
			Altmail:          entry.GetAttributeValue("altmail"),
			ID:               dcuID,
			Course:           entry.GetAttributeValue("course"),
			Year:             year,
			YearsPaid:        yearsPaid,
			UpdatedBy:        entry.GetAttributeValue("updatedBy"),
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
			Host:             entry.GetAttributeValues("host"),
			ShadowLastChange: shadow,
		})
	}
	return users, err
}
