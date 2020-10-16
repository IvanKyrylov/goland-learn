package taskone

import (
	"fmt"
	"log"
	"time"

	"github.com/beevik/ntp"
)

const server = "3.pool.ntp.org"

func HelloNow() time.Time {
	timeNtp, err := ntp.Time(server)
	if err != nil {
		log.Fatalf("Error: %s", err)
		return timeNtp
	}

	localTime := time.Now()

	fmt.Printf("Local time: %v\n", localTime.Round(time.Nanosecond).String())
	fmt.Printf("Exact time: %v\n", timeNtp.Round(time.Nanosecond).String())

	return timeNtp
}
