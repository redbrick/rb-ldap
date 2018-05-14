package rbuser

const expiredShell = "/usr/local/shells/expired"
const noLoginShell = "/usr/local/shells/disusered"
const defaultShell = "/usr/local/shells/shell"

// DisableUser disable an account from ldap
func (rb *RbLdap) DisableUser(user RbUser) error {
	user.LoginShell = noLoginShell
	return rb.Update(user)
}

// ExpireUser disable an account from ldap
func (rb *RbLdap) ExpireUser(user RbUser) error {
	user.LoginShell = expiredShell
	return rb.Update(user)
}

// RenableUser renable an account from ldap
func (rb *RbLdap) RenableUser(user RbUser) error {
	return rb.ResetShell(user)
}

// ResetShell reset a users shell
func (rb *RbLdap) ResetShell(user RbUser) error {
	user.LoginShell = defaultShell
	return rb.Update(user)
}
