package rbuser

import (
	"fmt"
	"time"

	ldap "gopkg.in/ldap.v2"
)

// Update a user in ldap
func (rb *RbLdap) Update(user RbUser) error {
	modification := ldap.NewModifyRequest(fmt.Sprintf("uid=%s,ou=ldap,o=redbrick", user.UID))
	now := time.Now()
	modification.Replace("cn", []string{user.CN})
	modification.Replace("altmail", []string{user.Altmail})
	modification.Replace("course", []string{user.Course})
	modification.Replace("year", []string{string(user.Year)})
	modification.Replace("updated", []string{now.Format(timeLayout)})
	modification.Replace("updatedBy", []string{user.UpdatedBy})
	modification.Replace("loginShell", []string{user.LoginShell})
	modification.Replace("birthday", []string{user.Birthday.Format(timeLayout)})
	return rb.Conn.Modify(modification)
}
