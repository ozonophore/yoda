package observer

type SchedulerObserver interface {
	BeforeJobExecution(jodId int)
}
