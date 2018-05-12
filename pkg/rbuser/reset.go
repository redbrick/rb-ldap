package rbuser

import (
	"fmt"

	ldap "gopkg.in/ldap.v2"
)

// ResetPasswd Reset A users password, takes users uid and boolean if to email them
func (rb *RbLdap) ResetPasswd(uid string) error {
	passwordModifyRequest := ldap.NewPasswordModifyRequest(fmt.Sprintf("uid=%s,ou=ldap,o=redbrick", uid), "", passwd(12))
	passwordModifyResponse, err := rb.Conn.PasswordModify(passwordModifyRequest)
	if err != nil {
		return err
	}
	user, err := rb.Search(fmt.Sprintf("(uid=%s)", uid))
	if err != nil {
		return err
	}
	user.UserPassword = passwordModifyResponse.GeneratedPassword
	return rb.mailAccountUpdate(user)
}
