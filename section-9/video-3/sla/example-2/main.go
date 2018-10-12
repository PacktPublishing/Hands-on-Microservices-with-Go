package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	timePeriodInDays := 365

	slaOurService := 0.999
	slaDBService := 0.999

	sla := slaOurService * slaDBService

	numSecsInPeriod := timePeriodInDays * 24 * 60 * 60
	uptime := int(sla * float64(numSecsInPeriod))
	downtime := numSecsInPeriod - uptime
	downtimeStr := strconv.Itoa(downtime) + "s"

	dur, _ := time.ParseDuration(downtimeStr)
	fmt.Println(dur)

}
