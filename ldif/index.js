#!/usr/bin/env node

/*
 * ldapsearch -D cn=root,ou=ldap,o=redbrick -xLLL \
 * -y /etc/ldap.secret \
 *  objectClass uid gidNumber > entry.ldif
 */

const fs = require('fs-extra');
const { isEmpty, isUndefined } = require('lodash');
const ldif = require('ldif');

const vhost = user =>
  `use VHost /storage/webtree/${user.username.charAt(0)}/${user.username} ${user.username} ${
    user.group
  } ${user.username}\n`;

const ldap = ldif.parseFile('./entry.ldif');
fs.appendFile(
  'user_vhost_list.conf',
  ldap.entries
    .map(convertUser)
    .filter(u => !isUndefined(u))
    .join('\n'),
);

function convertUser(entry) {
  const { attributes: { objectClass, uid, gidNumber } } = entry.toObject();
  const user = { username: uid };
  if (
    objectClass === 'club' ||
    objectClass === 'committe' ||
    objectClass === 'society' ||
    objectClass === 'associate' ||
    objectClass === 'member'
  ) {
    user.group = objectClass;
  }
  if (isEmpty(user.group)) user.group = getGroup(gidNumber);
  if (!isEmpty(user.username) && !isEmpty(user.group)) return vhost(user);
  return undefined;
}

function getGroup(gid) {
  switch (gid) {
    case 100:
      return 'committe';
    case 101:
      return 'society';
    case 102:
      return 'club';
    case 103:
      return 'member';
    case 105:
      return 'founder';
    case 107:
      return 'associat';
    case 31382:
      return 'dcu';
    case 109:
      return 'staff';
    case 1014:
      return 'projects';
    case 1017:
      return 'redbrick';
    case 1016:
      return 'intersoc';
    default:
      return 'member';
  }
}
