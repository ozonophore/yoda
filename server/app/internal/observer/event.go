package observer

type EventObserver interface {
	RunImmediately(jobID int)
}
