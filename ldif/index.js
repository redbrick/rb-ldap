#!/usr/bin/env node

/*
 * ldapsearch -D cn=root,ou=ldap,o=redbrick -xLLL \
 * -y /etc/ldap.secret \
 *  objectClass uid gidNumber > entry.ldif
 */

const fs = require('fs-extra');
const { isUndefined, isEmpty } = require('lodash');
const ldif = require('ldif');

const ldap = ldif.parseFile('./entry.ldif');
const users = [];

function getGroup({ gid }) {
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

(() => {
  ldap.entries.forEach(({ attributes }) => {
    if (!isUndefined(attributes)) {
      const user = {};
      let reserved = false;
      attributes.forEach(({ attribute, value }) => {
        if (attribute.attribute === 'uid') {
          user.username = value.value;
        }
        if (attribute.attribute === 'gidNumber') {
          user.gid = value.value;
        }
        if (
          value.value === 'club' ||
          value.value === 'committe' ||
          value.value === 'society' ||
          value.value === 'associate' ||
          value.value === 'member'
        ) {
          user.group = value.value;
        } else if (value.value === 'redbrick' || value.value === 'reserved') {
          reserved = true;
        }
      });
      if (!isEmpty(user) && !reserved) users.push(user);
    }
  });

  users.forEach(user => {
    if (isEmpty(user.group)) user.group = getGroup(user);
    if (!isEmpty(user.username) && !isEmpty(user.group)) {
      fs.appendFile(
        'user_vhost_list.conf',
        `use VHost /storage/webtree/${user.username.charAt(
          0,
        )}/${user.username} ${user.username} ${user.group} ${user.username}\n`,
      );
    }
  });
})();
