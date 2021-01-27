package cronjob

import (
	"errors"

	"github.com/robfig/cron/v3"
)

const (
	RetryTask = "retryTask"
)

var (
	AllJob = map[string]*JobInfo{
		RetryTask: nil,
	}
)

type JobInfo struct {
	Job     cron.Job
	Spec    string
	entryID cron.EntryID
}

func NewJobInfo(job cron.Job, spec string) *JobInfo {
	return &JobInfo{
		Spec: spec,
		Job:  job,
	}
}

func (j *JobInfo) RegisterCron(c *cron.Cron) error {
	if j.entryID > 0 {
		return errors.New("Job has been register cron")
	}
	var err error
	j.entryID, err = c.AddJob(j.Spec, j.Job)
	return err
}

func (j *JobInfo) RemoveCron(c *cron.Cron) {
	c.Remove(j.entryID)
	j.entryID = 0
}
