package job

type RetryTaskJob struct {
}

func (r *RetryTaskJob) Run() {
	// lock
	// get jobs and update excute time
	// unlock
	// async excute
}
