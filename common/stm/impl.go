package stm

import (
	"fmt"
	"sync"
)

type defaultSTM struct {
	sync.RWMutex

	init  State
	state State
	table map[State]Handler
}

func newDefaultSTM(init State, table map[State]Handler) *defaultSTM {
	return &defaultSTM{init: init, state: init, table: table}
}

func (d *defaultSTM) State() State {
	d.RLock()
	defer d.RUnlock()

	return d.state
}

func (d *defaultSTM) Reset() {
	d.Lock()
	defer d.Unlock()

	d.state = d.init
}

func (d *defaultSTM) Process(evt interface{}) (interface{}, error) {
	if h := d.table[d.state]; h != nil {
		return h(&defaultCtx{stm: d, msg: evt})
	} else {
		return nil, fmt.Errorf("state %d doesn't have a handler, it can't receive event: %v", d.state, evt)
	}
}

func (d *defaultSTM) setState(s State) error {
	d.Lock()
	defer d.Unlock()

	if _, ok := d.table[s]; ok {
		d.state = s
		return nil
	} else {
		return fmt.Errorf("state not registered: %d", s)
	}
}
