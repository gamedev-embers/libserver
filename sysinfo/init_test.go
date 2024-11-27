package sysinfo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShow(t *testing.T) {
	Show()
	// t.Fail()
}

func TestSysInfo(t *testing.T) {
	assert := assert.New(t)

	logs := Dump()
	assert.NotEmpty(logs)

	for _, line := range logs {
		t.Log(line)
	}
}
