package rbuser

import (
	"fmt"

	ldap "gopkg.in/ldap.v2"
)

// LdapConf server object used for connecting to server
type LdapConf struct {
	User     string
	Password string
	Host     string
	Port     int
}

// Connect to ldap database
func (conf *LdapConf) Connect() (*ldap.Conn, error) {
	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", conf.Host, conf.Port))
	if err != nil {
		return l, err
	}
	defer l.Close()

	err = l.Bind(conf.User, conf.Password)
	return l, err
}
