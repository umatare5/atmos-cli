package command

import (
	"github.com/urfave/cli/v2"
)

// RegisterCommands returns command structures
func RegisterCommands() []*cli.Command {
	command := []*cli.Command{}
	command = append(command, registerProfileCommand()...)
	command = append(command, registerDivelogCommand()...)
	return command
}
