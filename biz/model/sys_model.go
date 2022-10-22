package model

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

type SysInfo struct {
	CpuPercent float64                `json:"cpu_percent"`
	CpuInfo    *cpu.InfoStat          `json:"cpu_info"`
	MemInfo    *mem.VirtualMemoryStat `json:"mem_info"`
}

func (s *SysInfo) String() string {
	b, err := json.Marshal(*s)
	if err != nil {
		return fmt.Sprintf("%+v", *s)
	}
	var out bytes.Buffer
	err = json.Indent(&out, b, "", "    ")
	if err != nil {
		return fmt.Sprintf("%+v", *s)
	}
	return out.String()
}
