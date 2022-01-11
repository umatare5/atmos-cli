package command

import (
	"github.com/umatare5/atmos-cli/internal/domain"

	"github.com/urfave/cli/v2"
)

// registerAccessTokenFlag
func registerAccessTokenFlag() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    domain.AccessTokenFlagName,
			Usage:   domain.AccessTokenFlagUsage,
			Aliases: domain.AccessTokenFlagAliases,
			EnvVars: domain.AccessTokenFlagEnvVars,
		},
	}
}

// registerLimitFlag
func registerLimitFlag() []cli.Flag {
	return []cli.Flag{
		&cli.IntFlag{
			Name:    domain.LimitFlagName,
			Value:   domain.LimitFlagDefaultValue,
			Usage:   domain.LimitFlagUsage,
			Aliases: domain.LimitFlagAliases,
		},
	}
}

// registerCursorFlag
func registerCursorFlag() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    domain.CursorFlagName,
			Value:   domain.CursorFlagDefaultValue,
			Usage:   domain.CursorFlagUsage,
			Aliases: domain.CursorFlagAliases,
		},
	}
}

// registerStatisticsFlag
func registerStatisticsFlag() []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:    domain.StatisticsFlagName,
			Usage:   domain.StatisticsFlagUsage,
			Value:   domain.StatisticsFlagDefaultValue,
			Aliases: domain.StatisticsFlagAliases,
		},
	}
}

// registerUserIDFlag
func registerUserIDFlag() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    domain.UserIDFlagName,
			Usage:   domain.UserIDFlagUsage,
			Aliases: domain.UserIDFlagAliases,
			EnvVars: domain.UserIDFlagEnvVars,
		},
	}
}

// registerDivelogIDFlag
func registerDivelogIDFlag() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    domain.DivelogIDFlagName,
			Value:   domain.DivelogIDFlagDefaultValue,
			Usage:   domain.DivelogIDFlagUsage,
			Aliases: domain.DivelogIDFlagAliases,
		},
	}
}

// registerPrettyFormatFlag
func registerPrettyFormatFlag() []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:    domain.PrettyFormatFlagName,
			Usage:   domain.PrettyFormatFlagUsage,
			Value:   domain.PrettyFormatFlagDefaultValue,
			Aliases: domain.PrettyFormatFlagAliases,
		},
	}
}
