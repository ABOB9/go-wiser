package stm

type (
	State uint32

	Context interface {
		State() State
		Become(s State) error
		Event() interface{}
	}

	Handler func(ctx Context) (interface{}, error)

	StateMachine interface {
		State() State
		Process(evt interface{}) (interface{}, error)
		Reset()
	}

	Builder interface {
		SetInitState(State)
		RegState(State, Handler)
		Build() (StateMachine, error)
	}
)
