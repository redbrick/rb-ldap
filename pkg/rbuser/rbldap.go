package rbuser

// RbLdap Server object used for connecting to server
type RbLdap struct {
	*ldapConf
}

// NewRbLdap create ldap connection to Redbrick LDAP
func NewRbLdap(user, password, host string, port int) (*RbLdap, error) {
	conf := &ldapConf{user: user, password: password, host: host, port: port}
	rb := RbLdap{conf}
	rb.connect()
	return &rb, nil
}
