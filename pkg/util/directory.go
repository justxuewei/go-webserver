package util

import "os/user"

func GetHomeDirectory() string {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	return usr.HomeDir
}
