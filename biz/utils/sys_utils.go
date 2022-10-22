package utils

import (
	"time"

	"github.com/shirou/gopsutil/mem"

	"github.com/shirou/gopsutil/cpu"
)

// GetCpuInfo 获取CPU详细信息
func GetCpuInfo() (*cpu.InfoStat, error) {
	cpuInfos, err := cpu.Info()
	if err != nil {
		return nil, err
	}
	return &cpuInfos[0], nil
}

// GetCpuPercent 获取CPU使用率
func GetCpuPercent() (float64, error) {
	percent, err := cpu.Percent(time.Second, false)
	if err != nil {
		return 0, err
	}
	return percent[0], nil
}

func GetMemInfo() (*mem.VirtualMemoryStat, error) {
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}
	return memInfo, nil
}
