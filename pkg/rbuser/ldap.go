package rbuser

import (
	"fmt"

	ldap "gopkg.in/ldap.v2"
)

// LdapConf Server object used for connecting to server
type ldapConf struct {
	user     string
	password string
	host     string
	port     int
	Conn     *ldap.Conn
}

// Connect to ldap database
func (conf *ldapConf) connect() error {
	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", conf.host, conf.port))
	if err != nil {
		return err
	}
	conf.Conn = l
	defer conf.Conn.Close()
	return conf.Conn.Bind(conf.user, conf.password)
}
