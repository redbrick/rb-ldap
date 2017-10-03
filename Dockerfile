FROM node:8.6

# Create app directory
RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app

# Install app dependencies
COPY package.json /usr/src/app/
COPY yarn.lock /usr/src/app/

ADD . /usr/src/app

RUN yarn install

CMD ["node", "index.js", "ldapsearch -D cn=root,ou=ldap,o=redbrick -xLLL -y /etc/ldap.secret objectClass uid gidNumber > entry.ldif"]