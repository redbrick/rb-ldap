package rbuser

import (
	"fmt"

	ldap "gopkg.in/ldap.v2"
)

// LdapConf Server object used for connecting to server
type LdapConf struct {
	User     string
	Password string
	Host     string
	Port     int
	Conn     *ldap.Conn
}

// Connect to ldap database
func (conf *LdapConf) Connect() error {
	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", conf.Host, conf.Port))
	if err != nil {
		return err
	}
	defer l.Close()
	conf.Conn = l
	return l.Bind(conf.User, conf.Password)
}
