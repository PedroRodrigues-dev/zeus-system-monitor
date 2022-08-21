package core

import (
	"fmt"
	"time"
	"zeus/scripts"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

func Monitors(option int) {
	// In process variables
	timeLastNotification := time.Now()
	overload_counter := 0

	// In enviroment variables
	var percentLimit int
	var overloadCounterLimit int
	var overloadMessage string

	for {
		// Choose which percentage to use
		percentAvg := switcher(option, &percentLimit, &overloadCounterLimit, &overloadMessage)

		// Checks if percentage exceeded the threshold
		if percentAvg > percentLimit {
			overload_counter += 1
		}

		// Checks if the notification should be sent
		if overload_counter > overloadCounterLimit && time.Since(timeLastNotification) > time.Duration(scripts.ENV.TimeNotificationLimit*int(time.Second)) {
			timeLastNotification = time.Now()
			overload_counter = 0
			Notificator(fmt.Sprintf(overloadMessage, percentLimit))
		}

		// Avoid overload by pausing the goroutine for 0.1 seconds
		time.Sleep(time.Duration(1e+9))
	}
}

func switcher(option int, percentLimit *int, overloadCounterLimit *int, overloadMessage *string) int {
	// Define percentLimit, overloadCounterLimit and overloadMessage pointers and return percent now
	switch {
	// For CPU
	case option == 0:
		// Pointers
		*percentLimit = scripts.ENV.CpuPercentLimit
		*overloadCounterLimit = scripts.ENV.CpuOverloadCounterLimit
		*overloadMessage = scripts.ENV.CpuOverloadMessage

		// Require percents
		cpu_percents, _ := cpu.Percent(0, true)
		percents_sum := 0.0

		// Percents Avg
		for _, percent := range cpu_percents {
			percents_sum += percent
		}
		return int(percents_sum) / len(cpu_percents)

	// For memory
	case option == 1:
		// Pointers
		*percentLimit = scripts.ENV.MemoryPercentLimit
		*overloadCounterLimit = scripts.ENV.MemoryOverloadCounterLimit
		*overloadMessage = scripts.ENV.MemoryOverloadMessage

		// Percents
		memory, _ := mem.VirtualMemory()
		return int(memory.UsedPercent)

	// For disk
	case option == 2:
		// Pointers
		*percentLimit = scripts.ENV.DiskPercentLimit
		*overloadCounterLimit = scripts.ENV.DiskOverloadCounterLimit
		*overloadMessage = scripts.ENV.DiskOverloadMessage

		// Percents
		diskInfo, _ := disk.Usage("/")
		return int(diskInfo.UsedPercent)
	default:
		return 0
	}
}
