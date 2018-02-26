package rbuser

import (
	"fmt"

	ldap "gopkg.in/ldap.v2"
)

// RbLdap Server object used for connecting to server
type RbLdap struct {
	User     string
	Password string
	Host     string
	Port     int
	Conn     *ldap.Conn
}

// Connect to ldap database
func (conf *RbLdap) Connect() error {
	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", conf.Host, conf.Port))
	if err != nil {
		return err
	}
	defer l.Close()
	conf.Conn = l
	return l.Bind(conf.User, conf.Password)
}
