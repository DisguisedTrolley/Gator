package main

import "fmt"

type command struct {
	name string
	args []string
}

type commands struct {
	cmd map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	fn, ok := c.cmd[cmd.name]
	if ok {
		err := fn(s, cmd)
		if err != nil {
			return err
		}

		return nil
	}

	return fmt.Errorf("unknown command")
}

func (c *commands) register(name string, f func(*state, command) error) {
	if _, ok := c.cmd[name]; !ok {
		c.cmd[name] = f
	}
}
