# rb-ldap

[![Go Report Card](https://goreportcard.com/badge/github.com/redbrick/rb-ldap)](https://goreportcard.com/report/github.com/redbrick/rb-ldap)

Script to interact with Redbrick LDAP.

* query ldap for user info to generate apache template conf for user vhosts.
* Search for users in ldap

## Installation

```console
go get github.com/redbrick/rbldap/cmd/rb-ldap
```

## Run

```console
rb-ldap
```

Run `rb-ldap -h` to get a list of flags and commands.

```
NAME:
   rb-ldap - Command line interface for Redbrick LDAP

USAGE:
   rb-ldap [global options] command [command options] [arguments...]

COMMANDS:
     generate, g  generate list for uservhost macro
     search       Search ldap for user
     add, a       Add user to ldap
     help, h      Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --user value, -u value  ldap user, used for authentication (default: "cn=root,ou=ldap,o=redbrick")
   --dcu-user value        Active Directory user for DCU, used for authentication (default: "CN=rblookup,OU=Service Accounts,DC=ad,DC=dcu,DC=ie")
   --host value            ldap host to query (default: "ldap.internal")
   --dcu-host value        DCU Active Directory host to query (default: "ad.dcu.ie")
   --port value, -p value  Port for ldap host (default: 389)
   --dcu-port value        Port for DCU Active Directory host (default: 389)
   --password value        password for the ldap server [/etc/ldap.secret]
   --dcu-password value    password for the DCU ldap server [/etc/dcu_ldap.secret]
   --dry-run               output too console rather then file
   --help, -h              show help
```

## Notes

The conf from `rb-ldag g` will be output to the current dir.
