package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

func main() {
	timePeriodInDays := 365

	totalSla := 0.999 * math.Pow(0.9999, 4)

	numSecsInPeriod := timePeriodInDays * 24 * 60 * 60
	uptime := int(totalSla * float64(numSecsInPeriod))
	downtime := numSecsInPeriod - uptime
	downtimeStr := strconv.Itoa(downtime) + "s"

	dur, _ := time.ParseDuration(downtimeStr)
	fmt.Println(dur)

}
