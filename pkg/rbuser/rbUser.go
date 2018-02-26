package rbuser

import (
	"fmt"
	"html/template"
	"os"
	"time"
)

// RBUser is the redbric ldap user
type RBUser struct {
	UID              string
	UserType         string
	ObjectClass      string
	Newbie           bool   // New this year
	CN               string // Full name
	Altmail          string // Alternate email
	ID               int    // DCU ID number
	Course           string // DCU course code
	Year             int    // DCU course year number/code
	YearsPaid        int    // Number of years paid (integer)
	Updatedby        string // Username of user last to update
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
	ShadowLastChange time.Time
}

// Vhost reutrn apache  macro template
func (u *RBUser) Vhost() string {
	initial := []rune(u.UID)[0]
	return fmt.Sprintf("use VHost /storage/webtree/%s/%s %s %s %s", string(initial), u.UID, u.UID, u.ObjectClass, u.UID)
}

// PrettyPrint output user info to command line
func (u *RBUser) PrettyPrint() error {
	const output = `
	User Information
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
	{{ end }}{{ with .Updatedby }}updatedby: {{ . }}
	{{ end }}{{ with .Updated }}updated: {{ . }}
	{{ end }}{{ with .CreatedBy }}createdby: {{ . }}
	{{ end }}{{ with .Created }}created: {{ . }}
	{{ end }}{{ with .UIDNumber }} uidNumber: {{ . }}
	{{ end }}{{ with .GIDNumber }}gidNumber: {{ . }}
	{{ end }}{{ with .Gecos }}gecos: {{ . }}
	{{ end }}{{ with .LoginShell }}loginShell: {{ . }}{{ end }}
	{{ end }}{{ with .HomeDirectory }}homeDirectory: {{ . }}{{ end }}
	{{ end }}{{ with .UserPassword }}userPassword: {SHA}{{ . }}{{ end }}
	{{ end }}{{ with .Host }}host: {{ . }}
	{{ end }}{{ with .ShadowLastChange }}shadowLastChange: {{ . }}{{ end }}
`

	t := template.Must(template.New("user").Parse(output))
	return t.Execute(os.Stdout, u)
}
