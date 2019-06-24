package main

import (
	"fmt"
	"os"
	"strconv"
	"superquota/quota"
)

func main() {
	if len(os.Args) != 5 {
		fmt.Printf("Usage: %s <proj_id> <quota_bytes> <absolute_mountpoint_path> <ablosute_project_dir_path>\n", os.Args[0])
		os.Exit(1)
	}
	projectid := os.Args[1]
	quotaBytes := os.Args[2]
	mountPath := os.Args[3]
	path := os.Args[4]

	pid, err := strconv.Atoi(projectid)
	if err != nil {
		panic(err)
	}

	quotaInt, err := strconv.Atoi(quotaBytes)
	if err != nil {
		panic(err)
	}

	quotaControl, err := quota.NewControl(mountPath)
	if err != nil {
		panic(err)
	}

	err = quotaControl.SetQuota(uint32(pid), path, quota.Quota{Size: uint64(quotaInt)})
	if err != nil {
		panic(err)
	}
}
