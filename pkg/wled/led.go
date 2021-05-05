package wled

func (l *Led) SetState(state LedState) {
	l.lock.Lock()
	l.state = state
	l.lock.Unlock()
}
