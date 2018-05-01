package rbuser

import gomail "gopkg.in/gomail.v2"

// RbLdap Server object used for connecting to server
type RbLdap struct {
	*ldapConf
	mail *gomail.Dialer
}

// NewRbLdap create ldap connection to Redbrick LDAP
func NewRbLdap(user, password, host string, port int, smtp string) (*RbLdap, error) {
	conf := &ldapConf{user: user, password: password, host: host, port: port}
	d := gomail.Dialer{Host: smtp, Port: 587}
	rb := RbLdap{conf, &d}
	return &rb, rb.connect()
}
