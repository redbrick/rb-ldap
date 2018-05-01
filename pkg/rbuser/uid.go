package rbuser

import (
	"sort"
	"strconv"

	ldap "gopkg.in/ldap.v2"
)

func (rb *RbLdap) findAvailableUID() (int, error) {
	sr, err := rb.Conn.Search(ldap.NewSearchRequest(
		"ou=accounts,o=redbrick",
		ldap.ScopeSingleLevel, ldap.NeverDerefAliases,
		0, 0, false, "(&)",
		[]string{"uidNumber"}, nil,
	))
	if err != nil {
		return 0, err
	}
	UIDNumbers := make([]int, 0, len(sr.Entries))
	for _, entry := range sr.Entries {
		i, _ := strconv.Atoi(entry.GetAttributeValue("uidNumber"))
		UIDNumbers = append(UIDNumbers, i)
	}
	sort.Ints(UIDNumbers)
	return UIDNumbers[len(UIDNumbers)-1] + 1, nil
}
