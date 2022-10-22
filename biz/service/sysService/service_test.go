package sysService

import (
	"testing"
)

func TestGetSysInfo(t *testing.T) {
	if res, err := GetSysInfo(); err != nil {
		t.Errorf("getSysInfo failed, err: %v", err)
	} else {
		t.Logf("getSysInfo success, res: %v", res.String())
	}
}
