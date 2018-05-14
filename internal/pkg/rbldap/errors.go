package rbldap

import "errors"

var errNotImplemented = errors.New("dry-run not implemented")
var errUser404 = errors.New("User not found")
var errDuplicateUser = errors.New("User Already exists")
