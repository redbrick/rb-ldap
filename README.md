# userVhost

Script to query ldap for user info to generate apache template conf for user
vhosts.

## Installation

```console
go get github.com/redbrick/userVhost
```

## Run

```console
userVhost ./ldap.secret
```

The conf will be output to the current dir. Run `userVhost -h` to get a list of
flags
