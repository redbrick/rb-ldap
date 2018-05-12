package rbuser

import (
	"fmt"
	"time"

	ldap "gopkg.in/ldap.v2"
)

// Update a user in ldap
func (rb *RbLdap) Update(user RbUser) error {
	modification := ldap.NewModifyRequest(fmt.Sprintf("cn=%s,ou=ldap,o=redbrick", user.CN))
	now := time.Now()
	modification.Replace("cn", []string{user.CN})
	modification.Replace("altmail", []string{user.Altmail})
	modification.Replace("course", []string{user.Course})
	modification.Replace("year", []string{string(user.Year)})
	modification.Replace("updated", []string{now.Format("2006-01-02 15:04:00")})
	modification.Replace("updatedBy", []string{user.UpdatedBy})
	modification.Replace("loginShell", []string{user.LoginShell})
	modification.Replace("shadowlastchanged", []string{now.Format("2006-01-02 15:04:00")})
	modification.Replace("birthday", []string{user.Birthday.Format("2006-01-02 15:04:00")})
	return rb.Conn.Modify(modification)
}
