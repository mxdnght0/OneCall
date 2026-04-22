package onecall

import "sync"

type call struct {
	wg     sync.WaitGroup
	result any
	err    error
}
