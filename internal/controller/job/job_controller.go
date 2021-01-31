package job

type JobController struct {
	//audioService *audio.AudioService
}

func NewJobController() *JobController {
	j := &JobController{}
	//task.AddTaskRunner(task_constant.UploadAudioFile, task.NewTaskRunner())
	return j
}
