package main

import (
	"fmt"

	"github.com/lucasjo/porgex-go/usage/mem"
)

func main() {
	//dirpath := "/cgroup/memory/openshift"
	id := "55ee3a460f5106ab680000ca"

	//searchDir := filepath.Join(dirpath, "/", id)

	mData, _ := mem.GetAppMemory(id)

	fmt.Printf("MaxUsage - %v\n", mData.MaxUsage/1024/1024)
	fmt.Printf("Usage - %v\n", mData.Usage/1024/1024)
	fmt.Printf("Limit - %v\n", mData.Limit/1024/1024)

}
