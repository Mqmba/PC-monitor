package monitors

import (
	"context"
	"fmt"

	"github.com/shirou/gopsutil/v4/mem"
)

type MemMonitor struct {
}

func (mem *MemMonitor) Name() string {
	return "MEM"
}

func (m *MemMonitor) Check(ctx context.Context) (string, bool) {
	vmStat, err := mem.VirtualMemoryWithContext(ctx)
	if err != nil {
		return fmt.Sprintf("[Memory Monitor] could not retrieve Memory info: %v \n", err), false
	}

	// fmt.Printf("%+v", percent)
	value := fmt.Sprintf("%.2f%%", vmStat.UsedPercent)

	return value, vmStat.UsedPercent > 60
}
