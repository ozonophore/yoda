package observer

type EventObserver interface {
	RunImmediately(jobID int)
}

type EventObserverUpdateOrg interface {
	UpdateOrganizations() error
}
