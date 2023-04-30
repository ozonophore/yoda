package event

import "github.com/yoda/webapp/internal/observer"

var observers []observer.WSObserver

func AddObserver(observer observer.WSObserver) {
	observers = append(observers, observer)
}

func Notify(message []byte) {
	for _, observer := range observers {
		observer.BroadcastMessage(message)
	}
}
