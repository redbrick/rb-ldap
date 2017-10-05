#!/usr/bin/env node

const fs = require('fs-extra');
const { isUndefined, isEmpty } = require('lodash');
const ldap = require('ldapjs');

const client = ldap.createClient({
  url: `ldap://${process.env.LDAP_SEVRER || 'localhost'}`,
});

const opts = {
  attributes: ['objectClass', 'uid', 'gidNumber'],
};

(async () => {
  try {
    const secret = await fs.readFile(`${process.env.SECRET_FILE || './secret'}`, 'utf-8');
    client.bind('cn=root', secret, err => {
      if (err) throw err;
    });
    const users = await client.search(
      'cn=root,ou=ldap,o=redbrick',
      opts,
      (err, res) =>
        new Promise((resolve, reject) => {
          if (err) reject(err);
          const results = [];

          res.on('searchEntry', ({ object }) => {
            results.push(object);
          });

          res.on('error', ({ message }) => {
            console.error(`error: ${message}`);
          });

          res.on('end', ({ status }) => {
            console.log(`status: ${status}`);
            resolve(results);
          });
        }),
    );
    await fs.outputFile(
      `${process.env.OUTPUT_FILE || './user_vhost_list.conf'}`,
      convertUsers(users),
    );
  } catch (err) {
    console.errror(err);
    process.exit(1);
  }
})();

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

function convertUsers({ entries }) {
  const users = [];
  let vhosts = '';
  entries.forEach(({ attributes }) => {
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
      vhosts += `use VHost /storage/webtree/${user.username.charAt(
        0,
      )}/${user.username} ${user.username} ${user.group} ${user.username}\n`;
    }
  });
  return vhosts;
}
