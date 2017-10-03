# user-vhost-generator

Script to convert an ldif in to conf for uservhost

## To run

``` console
yarn
ldapsearch -D cn=root,ou=ldap,o=redbrick -xLLL -y /etc/ldap.secret objectClass uid gidNumber > entry.ldif
node .
```

## docker

First of all, you need to build a container:

`docker build -t user-vhost-generator`

Next, install project dependencies:

`docker run -v ${PWD}:/usr/src/app -t user-vhost-generator yarn install`