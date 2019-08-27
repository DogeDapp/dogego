package tasks

import (
	"dogego/modules"
	"fmt"
	"time"

	"github.com/robfig/cron"
)

var Cron *cron.Cron

func StartCronJobs(locked bool) {
	Cron = cron.New()

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
			Cron.AddFunc(item.Time, func() { PublishTask(item) })
		}
	}

	Cron.Start()

	fmt.Println("Cron Jobs started success.")
}
