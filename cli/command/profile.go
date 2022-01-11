package command

import (
	"github.com/umatare5/atmos-cli/internal/application"
	"github.com/umatare5/atmos-cli/internal/config"
	"github.com/umatare5/atmos-cli/internal/domain"
	"github.com/umatare5/atmos-cli/internal/framework"
	"github.com/umatare5/atmos-cli/internal/infrastructure"

	"github.com/urfave/cli/v2"
)

// registerProfileFlags returns command flags
func registerProfileFlags() []cli.Flag {
	flags := []cli.Flag{}
	flags = append(flags, registerAccessTokenFlag()...)
	flags = append(flags, registerUserIDFlag()...)
	flags = append(flags, registerStatisticsFlag()...)
	flags = append(flags, registerPrettyFormatFlag()...)
	return flags
}

// registerProfileCommand executes the command
func registerProfileCommand() []*cli.Command {
	return []*cli.Command{
		{
			Name:    domain.ProfileCommandName,
			Usage:   domain.ProfileCommandUsage,
			Aliases: domain.ProfileCommandAlias,
			Flags:   registerProfileFlags(),
			Action: func(ctx *cli.Context) error {
				c := config.New(ctx)
				r := infrastructure.New(&c)
				u := application.New(&c, &r)
				exec := framework.New(&c, &r, &u)

				if ctx.Bool(domain.StatisticsFlagName) {
					exec.GetUserStatistics(
						ctx.String(domain.UserIDFlagName),
						ctx.Bool(domain.PrettyFormatFlagName),
					)
					return nil
				}

				exec.GetUserProfile(
					ctx.String(domain.UserIDFlagName),
					ctx.Bool(domain.PrettyFormatFlagName),
				)
				return nil
			},
		},
	}
}
