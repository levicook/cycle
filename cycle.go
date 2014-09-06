package cycle

import "sync"

type StringCycle interface {
	Next() string
}

func Strings(strings ...string) StringCycle {
	if len(strings) == 0 {
		strings = []string{""}
	}

	c := new(stringCycle)
	c.Source = strings
	return c
}

type stringCycle struct {
	Source []string
	Index  int
	sync.Mutex
}

func (c *stringCycle) Next() string {
	v := c.Source[c.Index]

	c.Lock()
	c.Index = (c.Index + 1) % len(c.Source)
	c.Unlock()

	return v
}
