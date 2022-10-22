package sysService

import (
	"github.com/xhdd123321/whicinth-steganography-bd/biz/model"
	"github.com/xhdd123321/whicinth-steganography-bd/biz/utils"
)

// GetSysInfo 获取系统性能信息
func GetSysInfo() (*model.SysInfo, error) {
	cpuPercent, err := utils.GetCpuPercent()
	if err != nil {
		return nil, err
	}
	cpuInfo, err := utils.GetCpuInfo()
	if err != nil {
		return nil, err
	}

	memInfo, err := utils.GetMemInfo()
	if err != nil {
		return nil, err
	}
	return &model.SysInfo{
		CpuPercent: cpuPercent,
		CpuInfo:    cpuInfo,
		MemInfo:    memInfo,
	}, nil
}
