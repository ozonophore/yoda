package listener

type Listener interface {
	RunTask(jobId int)
}
