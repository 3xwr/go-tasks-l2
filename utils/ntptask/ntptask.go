package ntptask

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

var ntpServer = "0.beevik-ntp.pool.ntp.org"

func printTime() time.Time {
	time, err := ntp.Time(ntpServer)
	if err != nil {
		println(err)
		os.Exit(1)
	}

	fmt.Println("NTP time ", time)
	return time
}
