package wled

func (l *Led) SetState(state LedState) {
	l.lock.Lock()
	l.state = state
	l.lock.Unlock()
}

func (l *Led) GetState() (state LedState) {
	l.lock.Lock()
	state = l.state
	l.lock.Unlock()
	return
}
