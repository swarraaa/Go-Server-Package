package goServer

import (
	"log"

	"github.com/robfig/cron/v3"
)

func StartCronWithRobfig(schedule string, task func()) {
    c := cron.New(cron.WithSeconds()) 
    _, err := c.AddFunc(schedule, task)
    if err != nil {
        log.Fatalf("Error in adding cron job: %v", err)
    }
    c.Start()
}

func keepAliveTask() {
    log.Println("Keep-alive task executed by robfig/cron!")
}
