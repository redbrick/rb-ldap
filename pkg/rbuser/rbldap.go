package rbuser

import gomail "gopkg.in/gomail.v2"

// RbLdap Server object used for connecting to server
type RbLdap struct {
	*ldapConf
	Mail *gomail.Dialer
}

// NewRbLdap create ldap connection to Redbrick LDAP
func NewRbLdap(user, password, host string, port int, smtp string) (*RbLdap, error) {
	rb := RbLdap{
		&ldapConf{
			user:     user,
			password: password,
			host:     host,
			port:     port,
		},
		&gomail.Dialer{
			Host: smtp,
			Port: 587,
		},
	}
	return &rb, rb.connect()
}
