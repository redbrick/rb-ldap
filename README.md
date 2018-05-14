# rb-ldap

[![Go Report Card](https://goreportcard.com/badge/github.com/redbrick/rb-ldap)](https://goreportcard.com/report/github.com/redbrick/rb-ldap)

Script to interact with Redbrick LDAP.

* query ldap for user info to generate apache template conf for user vhosts.
* Search for users in ldap
* Create a new Redbrick user
* Renew a user
* convert a users usertype
* edit user info
* reset a user's password and shell
* disable and renable a user account
* remove and lock unpaid accounts
* produce ldap stats

## Installation

```console
go get github.com/redbrick/rbldap/cmd/rb-ldap
```

## Run

```console
rb-ldap
```

Run `rb-ldap -h` to get a list of flags and commands.

```console
$ rb-ldap --help
NAME:
   rb-ldap - Command line interface for Redbrick LDAP

USAGE:
   rb-ldap [global options] command [command options] [arguments...]

VERSION:
   0.6.0

AUTHOR:
   Cian Butler <butlerx@redbrick.dcu.ie>

COMMANDS:
     add, a            Add user to ldap
     disable, disuser  Disable a Users ldap account
     generate, g       generate list for uservhost macro
     renable, reuser   Renable a Users ldap account
     renew, r          renew a LDAP user
     reset             reset a users password
     reset-shell       reset a users shell
     search, s         Search ldap for user
     update, u, edit   Update a user in ldap
     help, h           Shows a list of commands or help for one command

   Batch Commands:
     alert-unpaid    Alert all unpaid users that their accounts will be disabled
     delete-unpaid   Delete all unpaid users accounts that are outside their grace period
     disable-unpaid  Diable all unpaid users accounts
     new-year        Decriment Years Paid of all users to 1

GLOBAL OPTIONS:
   --user value, -u value  ldap user, used for authentication (default: "cn=root,ou=ldap,o=redbrick")
   --dcu-user value        Active Directory user for DCU, used for authentication (default: "CN=rblookup,OU=Service Accounts,DC=ad,DC=dcu,DC=ie")
   --host value            ldap host to query (default: "ldap.internal")
   --dcu-host value        DCU Active Directory host to query (default: "ad.dcu.ie")
   --port value, -p value  Port for ldap host (default: 389)
   --dcu-port value        Port for DCU Active Directory host (default: 389)
   --password value        password for the ldap server [/etc/ldap.secret]
   --dcu-password value    password for the DCU ldap server [/etc/dcu_ldap.secret]
   --smtp value            smtp server to send email with [mailhost.redbrick.dcu.ie]
   --dry-run               output to console rather then file
   --help, -h              show help
   --version, -v           print the version
```

## Notes

The conf from `rb-ldag g` will be output to the current dir.
