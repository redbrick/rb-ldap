package rbuser

import (
	"fmt"
	"os/exec"
	"os/user"
	"strconv"
	"strings"
	"syscall"
)

// There is possibly a better way to interact with mailman but at the time of writing this appeared to be the easiest

func listAdd(list, email string) error {
	cmd := exec.Command("/var/lib/mailman/bin/add_members", "-r", "-", list)
	cmd.Stdin = strings.NewReader(fmt.Sprintf("%s\n", email))
	listUser, err := user.Lookup("list")
	if err != nil {
		return nil
	}
	uid, _ := strconv.ParseUint(listUser.Uid, 10, 32)
	gid, _ := strconv.ParseUint(listUser.Gid, 10, 32)
	cmd.SysProcAttr = &syscall.SysProcAttr{}
	// switch process uid so it is run by list user
	cmd.SysProcAttr.Credential = &syscall.Credential{Uid: uint32(uid), Gid: uint32(gid)}
	return cmd.Run()
}

func listDel(list, email string) error {
	cmd := exec.Command("/var/lib/mailman/bin/remove_members", list, email)
	cmd.Stdin = strings.NewReader(fmt.Sprintf("%s\n", email))
	listUser, err := user.Lookup("list")
	if err != nil {
		return nil
	}
	uid, _ := strconv.ParseUint(listUser.Uid, 10, 32)
	gid, _ := strconv.ParseUint(listUser.Gid, 10, 32)
	cmd.SysProcAttr = &syscall.SysProcAttr{}
	// switch process uid so it is run by list user
	cmd.SysProcAttr.Credential = &syscall.Credential{Uid: uint32(uid), Gid: uint32(gid)}
	return cmd.Run()
}
