package tasks

import (
	"dogego/modules"
	"fmt"
	"log"
	"time"

	"github.com/robfig/cron"
)

var Cron *cron.Cron

func StartCronJobs(locked bool) {
	Cron = cron.New()

	RegisterJobs()

	if !locked {
		if !modules.LockerModule.Lock("master", time.Minute*2) {
			Cron.AddFunc("@every 2m", CampaignMaster)
			Cron.Start()
			return
		}
	}

	Cron.AddFunc("@every 1m", ClifeMaster)

	for _, item := range modules.TasksModule {
		if item.Type == modules.TimeJob {
			log.Println(item)
			Cron.AddFunc(item.Time, func() { PublishTask(item) })
		}
	}

	Cron.Start()

	fmt.Println("Cron Jobs started success.")
}
