package commands

import (
	"fmt"
	"strings"

	"go.uber.org/multierr"
)

var (
	ErrInvalidChannel = fmt.Errorf("invalid channel number")
)

type CommandBuilder interface {
	Marshal() ([]byte, error)

	Channel(ch int) CommandBuilder
}

func NewCommandBuilder() CommandBuilder {
	return &command{
		sb: &strings.Builder{},
	}
}

type command struct {
	err error
	sb  *strings.Builder
}

func (c *command) Error() error {
	return c.err
}

func (c *command) Marshal() ([]byte, error) {
	if c.err != nil {
		return nil, c.err
	}

	return []byte(c.sb.String()), c.err
}

func (c *command) Channel(ch int) CommandBuilder {
	if ch < 1 || ch > 4 {
		c.err = multierr.Append(c.err, ErrInvalidChannel)
		return c
	}

	c.sb.WriteString(fmt.Sprintf(":CHAN%d", ch))
	return c
}
