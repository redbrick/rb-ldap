package rbuser

import (
	"fmt"

	ldap "gopkg.in/ldap.v2"
)

// DeleteUnPaid Delete unpaid accounts that have expired their grace period
func (rb *RbLdap) DeleteUnPaid() error {
	users, err := rb.SearchUsers("(&((yearspaid=-1))(|(usertype=member)(usertype=associate)(usertype=staff))")
	if err != nil {
		return err
	}
	for _, user := range users {
		if err := rb.DeleteUser(user); err != nil {
			return err
		}
	}
	return nil
}

// DisableUnPaid Disable unpaid accounts
func (rb *RbLdap) DisableUnPaid(admin string) error {
	users, err := rb.SearchUsers("(&((yearspaid=0))(|(usertype=member)(usertype=associate)(usertype=staff))")
	if err != nil {
		return err
	}
	for _, user := range users {
		user.UpdatedBy = admin
		if err := rb.ExpireUser(user); err != nil {
			return err
		}
	}
	return nil
}

// DeleteUser delete a users ldap account and home and web dir
func (rb *RbLdap) DeleteUser(user RbUser) error {
	if user.YearsPaid >= 0 {
		return nil
	}
	if err := rb.Conn.Del(
		ldap.NewDelRequest(fmt.Sprintf("uid=%s,ou=ldap,o=redbrick", user.UID), []ldap.Control{}),
	); err != nil {
		return err
	}
	if err := user.DelHomeDir(); err != nil {
		return err
	}
	return user.DelWebDir()
}
