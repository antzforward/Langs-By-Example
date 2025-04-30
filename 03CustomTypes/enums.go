package main

import "fmt"

type ServerState int //定义Enum底层是个int类型，嗯，看来可以改的，参考c#

const (
	StateIdle ServerState = iota //将可能的ServerState定义为常量，特殊关键字iota自动连续的常量值，只要首要的设置即可
	StateConnected
	StateError
	StateRetrying
)

var stateName = map[ServerState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

func (ss ServerState) String() string {
	return stateName[ss]
}

func main() {
	ns := transitions(StateIdle)
	fmt.Println(ns)

	ns2 := transitions(ns)
	fmt.Println(ns2)
}

func transitions(s ServerState) ServerState {

	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("unknown state: %s", s))
	}
}

/**go run .\enums.go
connected
idle
*/
