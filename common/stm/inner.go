package stm

type (
	innerStateMachine interface {
		StateMachine
		setState(s State) error
	}
)
