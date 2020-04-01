package stm

import (
	"fmt"
)

type defaultBuilder struct {
	init  State
	table map[State]Handler
}

func NewBuilder() *defaultBuilder {
	return &defaultBuilder{table: make(map[State]Handler)}
}

func (d *defaultBuilder) SetInitState(s State) {
	d.init = s
}

func (d *defaultBuilder) RegState(s State, h Handler) {
	d.table[s] = h
}

func (d *defaultBuilder) Build() (StateMachine, error) {
	if _, ok := d.table[d.init]; ok {
		return newDefaultSTM(d.init, d.table), nil
	} else {
		return nil, fmt.Errorf("init state or it's handler not set")
	}
}
