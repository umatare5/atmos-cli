package config

import (
	"log"

	"github.com/umatare5/atmos-cli/internal/domain"

	"github.com/jinzhu/configor"
	"github.com/urfave/cli/v2"
)

// Config struct
type Config struct {
	Server       string
	AccessToken  string
	Limit        int
	Cursor       string
	UserID       string
	DivelogID    string
	PrettyFormat bool
	Statistics   bool
}

// New returns Config struct
func New(ctx *cli.Context) Config {
	cfg := Config{
		Server:       ctx.String(domain.ServerFlagName),
		AccessToken:  ctx.String(domain.AccessTokenFlagName),
		Limit:        ctx.Int(domain.LimitFlagName),
		Cursor:       ctx.String(domain.CursorFlagName),
		UserID:       ctx.String(domain.UserIDFlagName),
		DivelogID:    ctx.String(domain.DivelogIDFlagName),
		PrettyFormat: ctx.Bool(domain.PrettyFormatFlagName),
		Statistics:   ctx.Bool(domain.StatisticsFlagName),
	}

	err := configor.New(&configor.Config{}).Load(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	err = checkArguments(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	return cfg
}

func checkArguments(cfg *Config) error {
	return nil
}
