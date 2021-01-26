package util

import "os"

func GetBaseDirFromEnvironment() string {
	dir, ok := os.LookupEnv("OSMROUTEDATADIR")
	if !ok {
		return "/run/shm/"
	}
	return dir
}
