package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

func main() {
	timePeriodInDays := 365
	estimatedDollarRevenue := 12000000

	totalSla := 0.99 * math.Pow(0.999, 4)

	numSecsInPeriod := timePeriodInDays * 24 * 60 * 60
	uptime := int(totalSla * float64(numSecsInPeriod))
	downtime := numSecsInPeriod - uptime

	revenuePerSecond := float64(estimatedDollarRevenue) / float64(numSecsInPeriod)

	costOfDowntime := float64(downtime) * revenuePerSecond

	downtimeStr := strconv.Itoa(downtime) + "s"
	dur, _ := time.ParseDuration(downtimeStr)
	fmt.Println("Downtime:", dur, " - Cost: $", int(costOfDowntime))

}
