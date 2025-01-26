package parser

import (
	"errors"
	"strings"
)

type CommandType string

const (
	SET CommandType = "SET"
	GET CommandType = "GET"
	DEL CommandType = "DEL"
)

type Command struct {
	Type CommandType
	Args []string
}

type Parser interface {
	Parse(input string) (*Command, error)
}

type SimpleParser struct{}

func NewParser() Parser {
	return &SimpleParser{}
}

func (p *SimpleParser) Parse(input string) (*Command, error) {
	tokens := strings.Fields(input)
	if len(tokens) == 0 {
		return nil, errors.New("empty input")
	}

	cmdType := CommandType(tokens[0])
	switch cmdType {
	case SET, GET, DEL:
		if err := validateArgs(cmdType, tokens[1:]); err != nil {
			return nil, err
		}
		return &Command{
			Type: cmdType,
			Args: tokens[1:],
		}, nil
	default:
		return nil, errors.New("unknown command")
	}
}

func validateArgs(cmdType CommandType, args []string) error {
	switch cmdType {
	case SET:
		if len(args) != 2 {
			return errors.New("SET command requires 2 arguments")
		}
	case GET, DEL:
		if len(args) != 1 {
			return errors.New(string(cmdType) + " command requires 1 argument")
		}
	}
	return nil
}
