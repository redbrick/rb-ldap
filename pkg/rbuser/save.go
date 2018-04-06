package rbuser

import (
	"fmt"
	"time"

	ldap "gopkg.in/ldap.v2"
)

// Add a user to ldap
func (rb *RbLdap) Add(user *RbUser) error {
	addition := ldap.NewAddRequest(fmt.Sprintf("cn=%s,ou=ldap,o=redbrick", user.CN))
	now := time.Now()
	/*
	   uidNumber: 102659
	   gidNumber: 103

	*/
	addition.Attribute("uid", []string{user.UID})
	addition.Attribute("usertype", []string{user.UserType})
	addition.Attribute("objectClass", []string{user.UserType, "posixAccount", "top", "shadowAccount"})
	addition.Attribute("newbie", []string{"true"})
	addition.Attribute("cn", []string{user.CN})
	addition.Attribute("altmail", []string{user.Altmail})
	addition.Attribute("id", []string{string(user.ID)})
	addition.Attribute("course", []string{user.Course})
	addition.Attribute("year", []string{string(user.Year)})
	addition.Attribute("yearspaid", []string{"1"})
	addition.Attribute("updated", []string{now.Format("2006-01-02 15:04:00")})
	addition.Attribute("updatedBy", []string{user.CreatedBy})
	addition.Attribute("created", []string{now.Format("2006-01-02 15:04:00")})
	addition.Attribute("createdBy", []string{user.CreatedBy})
	addition.Attribute("gecos", []string{user.CN})
	addition.Attribute("loginShell", []string{"/usr/local/shells/zsh"})
	addition.Attribute("homeDirectory", []string{"/home/" + user.UserType + "/" + string([]rune(user.UID)[0]) + "/" + user.UID})
	addition.Attribute("userPassword", []string{passwd(12)})
	addition.Attribute("host", user.Host)
	addition.Attribute("shadowlastchanged", []string{now.Format("2006-01-02 15:04:00")})
	addition.Attribute("birthday", []string{user.Birthday.Format("2006-01-02 15:04:00")})
	return rb.Conn.Add(addition)
}
