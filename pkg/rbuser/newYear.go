package rbuser

import (
	"fmt"
	"strconv"
	"time"

	ldap "gopkg.in/ldap.v2"
)

// NewYear to be run at the start of each year. Sets yearsPaid to 0 and set nood to false
func (rb *RbLdap) NewYear(admin string) error {
	users, err := rb.SearchUsers("(|(usertype=member)(usertype=associate)(usertype=staff))")
	if err != nil {
		return err
	}
	for _, user := range users {
		modification := ldap.NewModifyRequest(fmt.Sprintf("uid=%s,ou=ldap,o=redbrick", user.UID))
		now := time.Now()
		yearsPaid := strconv.Itoa(user.YearsPaid - 1)
		modification.Replace("newbie", []string{"FALSE"})
		modification.Replace("yearspaid", []string{yearsPaid})
		modification.Replace("updated", []string{now.Format("2006-01-02 15:04:00")})
		modification.Replace("updatedBy", []string{admin})
		if err := rb.Conn.Modify(modification); err != nil {
			return err
		}
	}
	return nil
}
