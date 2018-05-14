package rbuser

import (
	"fmt"
	"time"

	ldap "gopkg.in/ldap.v2"
)

// Renew a Users Account, variables yearsPaid will be set too 1
func (rb *RbLdap) Renew(user RbUser) error {
	modification := ldap.NewModifyRequest(fmt.Sprintf("uid=%s,ou=ldap,o=redbrick", user.UID))
	now := time.Now()
	modification.Replace("course", []string{user.Course})
	modification.Replace("yearspaid", []string{"1"})
	modification.Replace("year", []string{string(user.Year)})
	modification.Replace("updated", []string{now.Format(timeLayout)})
	modification.Replace("updatedBy", []string{user.UpdatedBy})
	if user.LoginShell == expiredShell || user.LoginShell == noLoginShell {
		modification.Replace("updatedBy", []string{defaultShell})
	}
	currentUser, err := rb.SearchUser(fmt.Sprintf("(&(uid=%s))", user.UID))
	if err != nil {
		return err
	}
	if currentUser.UserType != user.UserType {
		return rb.Conn.Modify(modification)
	}
	modification.Replace("objectClass", []string{user.UserType, "posixAccount", "top", "shadowAccount"})
	modification.Replace("homeDirectory", []string{user.HomeDirectory})
	if err := rb.Conn.Modify(modification); err != nil {
		return err
	}
	return rb.mailAccountUpdate(user)
}
