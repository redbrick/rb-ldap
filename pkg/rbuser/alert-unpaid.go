package rbuser

import (
	"bytes"
	"fmt"
	"html/template"

	gomail "gopkg.in/gomail.v2"
)

// AlertUnPaid emails all member, associate, & staff, with yearsPaid of 0 to tell them they are unpaid and remind them to renew
func (rb *RbLdap) AlertUnPaid() error {
	users, err := rb.SearchUsers("(&(|(yearspaid=0)(yearspaid=-1))(|(usertype=member)(usertype=associate)(usertype=staff))")
	if err != nil {
		return err
	}
	for _, user := range users {
		if err := rb.mailUnPaidWarning(user); err != nil {
			return err
		}
		fmt.Printf("%s has been warned\n", user.UID)
	}
	return nil
}

func (rb *RbLdap) mailUnPaidWarning(user RbUser) error {
	const email = `Hey there,
It's that time again to renew your Redbrick account!
Membership prices, as set by the SLC, are as follows:

Members      EUR 4
Associates   EUR 8
Staff        EUR 8
Guests       EUR 10

Note: if you have left DCU, you need to apply for associate membership.
You can pay in person, by lodging money into our account, electronic bank
transfer, or even PayPal! All the details you need are here:
https://www.redbrick.dcu.ie/help/joining/

Our bank details are:
a/c name: DCU Redbrick Society
IBAN: IE59BOFI90675027999600
BIC: BOFIIE2D
a/c number: 27999600
sort code: 90 - 67 - 50
Please Note!

{{ if eq .YearsPaid 0 }}
If you do not renew within the following month, your account will be disabled
Your account will remain on the system for a grace period of a year - you
just won't be able to login. So don't worry, it won't be deleted any time
soon! You can renew at any time during the year.
{{ else }}
If you do not renew within the following month, your account WILL BE
DELETED at the start of the new year. This is because you were not
recorded as having paid for last year and as such are nearing the end of
your one year 'grace' period to renew. Please make sure to renew as soon
as possible otherwise please contact us at: accounts@redbrick.dcu.ie.
{{ end }}

If in fact you have renewed and have received this email in error, it is
important you let us know. Just reply to this email and tell us how and
when you renewed and we'll sort it out.
For your information, your current Redbrick account details are:

{{- with .UID -}}username: {{ . }}
{{- end -}}{{- with .UserType -}}account type: {{ . }}
{{- end -}}{{- with .CN -}}name: {{ . }}
{{- end -}}{{- with .Altmail -}}alternative email: {{ . }}
{{- end -}}{{- with .ID -}}id number: {{ . }}
{{- end -}}{{- with .Course -}}course: {{ . }}
{{- end -}}{{- with .Year -}}year: {{ . }}{{ end }}

If any of the above details are wrong, please correct them when you
renew!
- Redbrick Admin Team
`
	t := template.Must(template.New("unpaid").Parse(email))
	m := gomail.NewMessage()
	m.SetHeader("From", "accounts@redbrick.dcu.ie")
	m.SetHeader("To", fmt.Sprintf("%s@redbrick.dcu.ie", user.UID))
	m.SetHeader("Cc", user.Altmail)
	m.SetHeader("Subject", "Time to renew your Redbrick account!")
	var tpl bytes.Buffer
	if err := t.Execute(&tpl, user); err != nil {
		return err
	}
	m.SetBody("text/plain", tpl.String())

	return rb.Mail.DialAndSend(m)
}
