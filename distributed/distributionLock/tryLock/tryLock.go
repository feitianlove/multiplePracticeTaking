package tryLock

type Lock struct {
	C chan struct{}
}

func NewLock() Lock {
	var l Lock
	l.C = make(chan struct{}, 1)
	l.C <- struct{}{}
	return l
}

func (l Lock) Lock() bool {
	lockResult := false
	select {
	case <-l.C:
		lockResult = true
	default:
	}
	return lockResult
}

func (l Lock) UnLock() {
	l.C <- struct{}{}
}
