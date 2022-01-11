package cli

import (
	"github.com/umatare5/atmos-cli/internal/domain"

	"github.com/urfave/cli/v2"
)

// registerServerFlag
func registerServerFlag() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    domain.ServerFlagName,
			Usage:   domain.ServerFlagUsage,
			Aliases: domain.ServerFlagAlias,
			EnvVars: domain.ServerFlagEnvVars,
		},
	}
}
