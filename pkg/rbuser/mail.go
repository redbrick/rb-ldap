package rbuser

import (
	"bytes"
	"html/template"

	gomail "gopkg.in/gomail.v2"
)

func (user *RbUser) mailAccountUpdate() error {
	const email = `
{{ if .Newbie }}
Welcome to Redbrick, the DCU Networking Society! Thank you for joining.
{{ else }}
Welcome back to Redbrick, the DCU Networking Society! Thank you for renewing.


Your Redbrick Account details are:

{{ with .UID }}username: {{ . }}
{{ end }}{{ with .UserPassword }}password: {{ . }}
{{ end }}{{ with .UserType }}account type: {{ . }}
{{ end }}{{ with .CN }}name: {{ . }}
{{ end }}{{ with .ID }}id number: {{ . }}
{{ end }}{{ with .Course }}course: {{ . }}
{{ end }}{{ with .Year }}year: {{ . }}{{ end }}

-------------------------------------------------------------------------------
your Redbrick webpage: https://{{ .UID }}.redbrick.dcu.ie/
your Redbrick email: {{ .UID -}}@redbrick.dcu.ie
You can find out more about our services at:
https://www.redbrick.dcu.ie/about/welcome

We recommend that you change your password as soon as you login.
Problems with your password or wish to change your username? Contact:
admin-request@redbrick.dcu.ie
Problems using Redbrick in general or not sure what to do? Contact:
helpdesk-request@redbrick.dcu.ie
Have fun!

- Redbrick Admin Team`
	t := template.Must(template.New("user").Parse(email))
	m := gomail.NewMessage()
	m.SetHeader("From", "admin-request@redbrick.dcu.ie")
	m.SetHeader("To", user.Altmail)
	m.SetHeader("Subject", "Welcome to Redbrick! - Your Account Details")
	var tpl bytes.Buffer
	if err := t.Execute(&tpl, user); err != nil {
		return err
	}
	m.SetBody("text/plain", tpl.String())

	d := gomail.Dialer{Host: "mailhost.redbrick.dcu.ie", Port: 587}
	return d.DialAndSend(m)
}
