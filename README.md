# user-vhost-generator

Script to convert an ldif in to conf for uservhost

## To run

``` console
yarn
ldapsearch -D cn=root,ou=ldap,o=redbrick -xLLL -y /etc/ldap.secret objectClass uid gidNumber > entry.ldif
node .
```

## docker

Build and run container with code:

`docker-compose -p uvg up --build`