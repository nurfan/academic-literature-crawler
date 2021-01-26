package workerpool

// JobQueue A buffered channel that we can send work requests on.
var JobQueue chan Job

// ExecutorInterface for job pattern function
type ExecutorInterface interface {
	Handle() error
}

// Job represents the job to be run
type Job struct {
	Executor ExecutorInterface
}

// SetExecutor setter job to job.executor
func (j *Job) SetExecutor(exec ExecutorInterface) {
	j.Executor = exec
}
