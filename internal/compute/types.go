package compute

type CommandType string

const (
	SET CommandType = "SET"
	GET CommandType = "GET"
	DEL CommandType = "DEL"
)

type Command struct {
	Type      CommandType
	Arguments []string
}

type Parser interface {
	Parse(input string) (*Command, error)
}

type Computer interface {
	Execute(cmd *Command) (string, error)
}
