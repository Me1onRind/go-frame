package cronjob

import (
	"github.com/robfig/cron/v3"
)

type CronJob struct {
	cron *cron.Cron
}

func NewCronJob() *CronJob {
	cron := cron.New()
	c := &CronJob{
		cron: cron,
	}
	return c
}
