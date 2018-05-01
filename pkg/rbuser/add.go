package rbuser

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"time"

	ldap "gopkg.in/ldap.v2"
)

// Add a user to ldap
func (rb *RbLdap) Add(user *RbUser) error {
	addition := ldap.NewAddRequest(fmt.Sprintf("cn=%s,ou=ldap,o=redbrick", user.CN))
	now := time.Now()
	uidNumber, err := rb.findAvailableUID()
	if err != nil {
		return err
	}
	gidNumber := groupToGID(user.UserType)
	addition.Attribute("gidNumber", []string{string(gidNumber)})
	addition.Attribute("uidNumber", []string{string(uidNumber)})
	addition.Attribute("uid", []string{user.UID})
	addition.Attribute("usertype", []string{user.UserType})
	addition.Attribute("objectClass", []string{user.UserType, "posixAccount", "top", "shadowAccount"})
	addition.Attribute("newbie", []string{"true"})
	addition.Attribute("cn", []string{user.CN})
	addition.Attribute("altmail", []string{user.Altmail})
	addition.Attribute("id", []string{string(user.ID)})
	addition.Attribute("course", []string{user.Course})
	addition.Attribute("year", []string{string(user.Year)})
	addition.Attribute("yearspaid", []string{"1"})
	addition.Attribute("updated", []string{now.Format("2006-01-02 15:04:00")})
	addition.Attribute("updatedBy", []string{user.CreatedBy})
	addition.Attribute("created", []string{now.Format("2006-01-02 15:04:00")})
	addition.Attribute("createdBy", []string{user.CreatedBy})
	addition.Attribute("gecos", []string{user.CN})
	addition.Attribute("loginShell", []string{"/usr/local/shells/zsh"})
	addition.Attribute("homeDirectory", []string{"/home/" + user.UserType + "/" + string([]rune(user.UID)[0]) + "/" + user.UID})
	addition.Attribute("userPassword", []string{passwd(12)})
	addition.Attribute("host", user.Host)
	addition.Attribute("shadowlastchanged", []string{now.Format("2006-01-02 15:04:00")})
	addition.Attribute("birthday", []string{user.Birthday.Format("2006-01-02 15:04:00")})
	if err := createHome(uidNumber, gidNumber, user); err != nil {
		return err
	}
	if err := createWebDir(uidNumber, gidNumber, user); err != nil {
		return err
	}
	if err := linkPublicHTML(user); err != nil {
		return err
	}
	return rb.Conn.Add(addition)
}

func (rb *RbLdap) findAvailableUID() (int, error) {
	sr, err := rb.Conn.Search(ldap.NewSearchRequest(
		"ou=accounts,o=redbrick",
		ldap.ScopeSingleLevel, ldap.NeverDerefAliases,
		0, 0, false, "(&())",
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

func createHome(uid, gid int, user *RbUser) error {
	folder := "/home/" + user.UserType + "/" + string([]rune(user.UID)[0]) + "/" + user.UID
	if err := os.MkdirAll(folder, os.ModePerm); err != nil {
		return err
	}
	return os.Chown(folder, uid, gid)
}

func createWebDir(uid, gid int, user *RbUser) error {
	folder := "/webtree/" + string([]rune(user.UID)[0]) + "/" + user.UID
	if err := os.MkdirAll(folder, os.ModePerm); err != nil {
		return err
	}
	return os.Chown(folder, uid, gid)
}

func linkPublicHTML(user *RbUser) error {
	return os.Symlink("/webtree/"+string([]rune(user.UID)[0])+"/"+user.UID, "/home/"+user.UserType+"/"+string([]rune(user.UID)[0])+"/"+user.UID+"/public_html")
}
