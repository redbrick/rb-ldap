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
	modification.Replace("updated", []string{now.Format("2006-01-02 15:04:00")})
	modification.Replace("updatedBy", []string{user.UpdatedBy})
	currentUser, err := rb.Search(fmt.Sprintf("(&(uid=%s))", user.UID))
	if err != nil {
		return err
	}
	if currentUser.UserType != user.UserType {
		return rb.Conn.Modify(modification)
	}
	modification.Replace("usertype", []string{user.UserType})
	modification.Replace("homeDirectory", []string{user.HomeDirectory})
	if err := rb.Conn.Modify(modification); err != nil {
		return err
	}
	return rb.mailAccountUpdate(user)
}
