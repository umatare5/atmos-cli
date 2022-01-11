package command

import (
	"github.com/umatare5/atmos-cli/internal/application"
	"github.com/umatare5/atmos-cli/internal/config"
	"github.com/umatare5/atmos-cli/internal/domain"
	"github.com/umatare5/atmos-cli/internal/framework"
	"github.com/umatare5/atmos-cli/internal/infrastructure"

	"github.com/urfave/cli/v2"
)

// registerDivelogFlags returns command flags
func registerDivelogFlags() []cli.Flag {
	flags := []cli.Flag{}
	flags = append(flags, registerAccessTokenFlag()...)
	flags = append(flags, registerDivelogIDFlag()...)
	flags = append(flags, registerLimitFlag()...)
	flags = append(flags, registerCursorFlag()...)
	flags = append(flags, registerPrettyFormatFlag()...)
	return flags
}

// registerDivelogCommand executes the command
func registerDivelogCommand() []*cli.Command {
	return []*cli.Command{
		{
			Name:    domain.DivelogCommandName,
			Usage:   domain.DivelogCommandUsage,
			Aliases: domain.DivelogCommandAlias,
			Flags:   registerDivelogFlags(),
			Action: func(ctx *cli.Context) error {
				c := config.New(ctx)
				r := infrastructure.New(&c)
				u := application.New(&c, &r)
				exec := framework.New(&c, &r, &u)

				if ctx.String(domain.DivelogIDFlagName) != domain.DivelogIDFlagDefaultValue {
					exec.GetDivelog(
						ctx.String(domain.DivelogIDFlagName),
						ctx.Bool(domain.PrettyFormatFlagName),
					)
					return nil
				}
				if ctx.Bool(domain.PrettyFormatFlagName) {
					exec.GetDivelogsPretty()
					return nil
				}

				exec.GetDivelogs()
				return nil
			},
		},
	}
}
