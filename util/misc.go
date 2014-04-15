package util

import "os/user"

func ApplicationPath() string {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	return usr.HomeDir + "/.storekeeper"
}
