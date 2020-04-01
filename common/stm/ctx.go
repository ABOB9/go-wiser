package stm

type defaultCtx struct {
	stm innerStateMachine
	msg interface{}
}

func (d *defaultCtx) State() State {
	return d.stm.State()
}

func (d *defaultCtx) Become(s State) error {
	return d.stm.setState(s)
}

func (d *defaultCtx) Event() interface{} {
	return d.msg
}
