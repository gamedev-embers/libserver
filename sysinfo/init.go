package sysinfo

import (
	"fmt"
	"log"
	"os"
	"runtime"
)

var info *SysInfo

func Show() {
	for _, line := range Dump() {
		log.Println(line)
	}

	for _, line := range DumpMem() {
		log.Println(line)
	}
}

func Dump() []string {
	obj := Load()
	return []string{
		fmt.Sprintf("GODEBUG      = %s", os.Getenv("GODEBUG")),
		fmt.Sprintf("limit.nofile = %+v", obj.MaxFD),
		fmt.Sprintf("runtime.NumCPU = %d", runtime.NumCPU()),
	}
}

func Load() *SysInfo {
	if info == nil {
		obj, err := newSysInfo()
		if err != nil {
			panic(err)
		}
		info = obj
	}
	return info
}
