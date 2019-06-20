package main

import "os"

import "strconv"
import "superquota/quota"

func main() {
	projectid := os.Args[1]
	path := os.Args[2]

	pid, err := strconv.Atoi(projectid)
	if err != nil {
		panic(err)
	}

	err = quota.SetProjectID(path, uint32(pid))
	if err != nil {
		panic(err)
	}
}
