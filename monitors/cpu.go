package monitors

import (
	"context"
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
)

type CPUMonitor struct {
}

func (cpu *CPUMonitor) Name() string {
	return "CPU"
}

func (m *CPUMonitor) Check(ctx context.Context) (string, bool) {
	cpuStat, err := cpu.PercentWithContext(ctx, 1*time.Second, false)
	if err != nil && len(cpuStat) == 0 {
		return fmt.Sprintf("[CPU Monitor] could not retrieve CPU info: %v \n", err), false
	}

	// fmt.Printf("%+v", cpuStat)
	value := fmt.Sprintf("%.2f%%", cpuStat[0])

	return value, cpuStat[0] > 60
}
