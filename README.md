# rb-ldap

[![Go Report Card](https://goreportcard.com/badge/github.com/redbrick/rbldap)](https://goreportcard.com/report/github.com/redbrick/rbldap)

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

## Notes

The conf from `rb-ldag g` will be output to the current dir.
