package monitors

import (
	"context"
	"fmt"
	"runtime"

	"github.com/shirou/gopsutil/disk"
)

type DiskMonitor struct {
}

func (cpu *DiskMonitor) Name() string {
	return "Disk"
}

func (m *DiskMonitor) Check(ctx context.Context) (string, bool) {
	path := "/"
	if runtime.GOOS == "Windows" {
		path = "C:"
	}
	diskStat, err := disk.UsageWithContext(ctx, path)
	if err != nil {
		return fmt.Sprintf("[Disk Monitor] could not retrieve Disk info: %v \n", err), false
	}

	// fmt.Printf("%+v", cpuStat)
	value := fmt.Sprintf("%.2f%% used", diskStat.UsedPercent)

	return value, diskStat.UsedPercent > 60
}
