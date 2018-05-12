package rbuser

import "os"

// CreateHome Create a users home dir and chown it to them
func (user *RbUser) CreateHome() error {
	if err := os.MkdirAll(user.HomeDirectory, os.ModePerm); err != nil {
		return err
	}
	return os.Chown(user.HomeDirectory, user.UIDNumber, user.GidNumber)
}

// CreateWebDir Create a users Web dir and chown it to them
func (user *RbUser) CreateWebDir() error {
	folder := "/webtree/" + string([]rune(user.UID)[0]) + "/" + user.UID
	if err := os.MkdirAll(folder, os.ModePerm); err != nil {
		return err
	}
	return os.Chown(folder, user.UIDNumber, user.GidNumber)
}

// LinkPublicHTML Link a users Webdir to their home dir
func (user *RbUser) LinkPublicHTML() error {
	return os.Symlink("/webtree/"+string([]rune(user.UID)[0])+"/"+user.UID, user.HomeDirectory+"/public_html")
}
