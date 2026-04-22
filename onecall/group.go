package onecall

import "sync"

type Group struct {
	mu    sync.Mutex
	calls map[string]*call
}

func NewGroup() *Group {
	return &Group{
		calls: make(map[string]*call),
	}
}

func (g *Group) Do(key string, f func() (any, error)) (any, error) {
	g.mu.Lock()
	c, found := g.calls[key]

	if found {
		g.mu.Unlock()
		c.wg.Wait()
		return c.result, c.err
	}

	c = &call{
		wg: sync.WaitGroup{},
	}
	c.wg.Add(1)
	g.calls[key] = c
	g.mu.Unlock()

	c.result, c.err = f()

	c.wg.Done()
	delete(g.calls, key)

	return c.result, c.err
}
