package cli

import (
	"github.com/umatare5/atmos-cli/cli/command"

	"errors"
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

// Preset parameters for this command
const (
	cmdName      string = "atmos"
	cmdUsage     string = "Client for ATMOS Platform"
	cmdUsageText string = "atmos COMMAND [options...]"
	cmdVersion   string = "0.1.0"
)

// registerGlobalFlags returns global flags
func registerGlobalFlags() []cli.Flag {
	flags := []cli.Flag{}
	flags = append(flags, registerServerFlag()...)
	return flags
}

// Start is a entrypoint of this command
func Start() {
	cmd := &cli.App{
		Name:      cmdName,
		HelpName:  cmdName,
		Usage:     cmdUsage,
		UsageText: cmdUsageText,
		Version:   cmdVersion,
		Flags:     registerGlobalFlags(),
		Commands:  command.RegisterCommands(),
		Action: func(ctx *cli.Context) error {
			err := cli.ShowAppHelp(ctx)
			if err != nil {
				return errors.New("Failed to invoke the help")
			}
			return nil
		},
	}

	err := cmd.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		return
	}
}
