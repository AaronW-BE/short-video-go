package utils

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"time"
)

type SystemInfo struct {
	Cpu struct {
		Cores   int
		Load    *load.AvgStat
		Percent []float64
	}
	Mem *mem.VirtualMemoryStat
}

func NewSysInfo() SystemInfo {
	cpuCount, err := cpu.Counts(true)
	if err != nil {
		cpuCount = 0
	}

	percent, _ := cpu.Percent(time.Second, true)

	memState, err := mem.VirtualMemory()

	loadInfo, _ := load.Avg()
	return SystemInfo{
		Cpu: struct {
			Cores   int
			Load    *load.AvgStat
			Percent []float64
		}{
			Cores:   cpuCount,
			Load:    loadInfo,
			Percent: percent,
		},
		Mem: memState,
	}
}
