package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	timePeriodInDays := 365
	sla := 0.99
	numSecsInPeriod := timePeriodInDays * 24 * 60 * 60
	uptime := int(sla * float64(numSecsInPeriod))
	downtime := numSecsInPeriod - uptime
	downtimeStr := strconv.Itoa(downtime) + "s"

	dur, _ := time.ParseDuration(downtimeStr)
	fmt.Println(dur)

}
