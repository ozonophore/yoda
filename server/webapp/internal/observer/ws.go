package observer

type WSObserver interface {
	BroadcastMessage(message []byte)
}
