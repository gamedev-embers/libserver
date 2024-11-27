package sysinfo

import (
	"runtime"

	"github.com/gamedev-embers/libserver/humanize"
)

func DumpMem() []string {
	m := MemUsage()
	return []string{
		"TotalAlloc = " + humanize.Size(m.TotalAlloc),
		"Alloc = " + humanize.Size(m.Alloc),
		"Sys = " + humanize.Size(m.Sys),
		"Lookups = " + humanize.Number(m.Lookups),
		"Mallocs = " + humanize.Number(m.Mallocs),
		"Frees = " + humanize.Number(m.Frees),
		"HeapAlloc = " + humanize.Size(m.HeapAlloc),
		"HeapSys = " + humanize.Size(m.HeapSys),
		"HeapIdle = " + humanize.Size(m.HeapIdle),
		"HeapInuse = " + humanize.Size(m.HeapInuse),
		"HeapReleased = " + humanize.Size(m.HeapReleased),
		"HeapObjects = " + humanize.Number(m.HeapObjects),
		"StackInuse = " + humanize.Size(m.StackInuse),
		"StackSys = " + humanize.Size(m.StackSys),
		"NextGC = " + humanize.TimeStamp(m.NextGC),
		"LastGC = " + humanize.TimeStamp(m.LastGC),
		"NumGC = " + humanize.Number(uint64(m.NumGC)),
	}
}

// MemUsage returns the memory usage statistics
// see: https://golang.org/pkg/runtime/#MemStats
func MemUsage() *runtime.MemStats {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return &m
}
