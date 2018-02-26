# rb-ldap

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

## Notes

The conf will be output to the current dir.

Run `userVhost -h` to get a list of flags and commands.
