package rbuser

import (
	"fmt"
	"html/template"
	"os"
	"time"
)

// RbUser is the redbric ldap user
type RbUser struct {
	UID              string
	UserType         string
	ObjectClass      []string
	Newbie           bool   // New this year
	CN               string // Full name
	Altmail          string // Alternate email
	ID               int    // DCU ID number
	Course           string // DCU course code
	Year             int    // DCU course year number/code
	YearsPaid        int    // Number of years paid (integer)
	UpdatedBy        string // Username of user last to update
	Updated          time.Time
	CreatedBy        string // Username of user that created them
	Created          time.Time
	Birthday         time.Time
	UIDNumber        int
	GidNumber        int
	Gecos            string
	LoginShell       string
	HomeDirectory    string
	UserPassword     string   // Crypted password.
	Host             []string // List of hosts.
	ShadowLastChange int
}

// Vhost reutrn apache  macro template
func (u *RbUser) Vhost() string {
	return fmt.Sprintf("use VHost /storage/webtree/%s/%s %s %s %s", string([]rune(u.UID)[0]), u.UID, u.UID, u.UserType, u.UID)
}

// PrettyPrint output user info to command line
func (u *RbUser) PrettyPrint() error {
	const output = `User Information
================
{{ with .UID }}uid: {{ . }}
{{ end }}{{ with .UserType }}usertype: {{ . }}
{{ end }}{{ with .ObjectClass }}objectClass: {{ . }}
{{ end }}{{ with .Newbie }}newbie: {{ . }}
{{ end }}{{ with .CN }}cn: {{ . }}
{{ end }}{{ with .Altmail }}altmail: {{ . }}
{{ end }}{{ with .ID }}id: {{ . }}
{{ end }}{{ with .Course }}course: {{ . }}
{{ end }}{{ with .Year }}year: {{ . }}
{{ end }}{{ with .YearsPaid }}yearsPaid: {{ . }}
{{ end }}{{ with .UpdatedBy }}updatedBy: {{ . }}
{{ end }}{{ with .Updated.Format "2006-01-02 15:04:05" }}updated: {{ . }}
{{ end }}{{ with .CreatedBy }}createdby: {{ . }}
{{ end }}{{ with .Created.Format "2006-01-02 15:04:05" }}created: {{ . }}
{{ end }}{{ with .Birthday.Format "2006-01-02 15:04:05" }}birthday: {{ . }}
{{ end }}{{ with .UIDNumber }}uidNumber: {{ . }}
{{ end }}{{ with .GidNumber }}gidNumber: {{ . }}
{{ end }}{{ with .Gecos }}gecos: {{ . }}
{{ end }}{{ with .LoginShell }}loginShell: {{ . }}
{{ end }}{{ with .HomeDirectory }}homeDirectory: {{ . }}
{{ end }}{{ with .UserPassword }}userPassword: {{ . }}
{{ end }}{{ with .Host }}host: {{ . }}
{{ end }}{{ with .ShadowLastChange }}shadowLastChange: {{ . }}{{ end }}
`

	t := template.Must(template.New("user").Parse(output))
	return t.Execute(os.Stdout, u)
}
