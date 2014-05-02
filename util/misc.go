package util

import "os/user"

func ApplicationPath() string {
	usr, _ := user.Current()
	return usr.HomeDir + "/.storekeeper"
}
