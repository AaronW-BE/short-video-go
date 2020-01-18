package utils

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

type SystemInfo struct {
	CpuAmount int
	Mem       *mem.VirtualMemoryStat
}

func NewSysInfo() SystemInfo {
	cpuCount, err := cpu.Counts(true)
	if err != nil {
		cpuCount = 0
	}

	memState, err := mem.VirtualMemory()

	return SystemInfo{
		CpuAmount: cpuCount,
		Mem:       memState,
	}
}
