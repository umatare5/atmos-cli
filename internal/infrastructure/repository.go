package infrastructure

import (
	"github.com/umatare5/atmos-cli/internal/config"

	atmosRepository "github.com/umatare5/atmos-cli/internal/infrastructure/atmos"
)

// Repository struct
type Repository struct {
	Config *config.Config
}

// New returns Repository struct
func New(c *config.Config) Repository {
	return Repository{
		Config: c,
	}
}

// InvokeAtmosRepository returns new AtmosRepository
func (r *Repository) InvokeAtmosRepository() *atmosRepository.Repository {
	return &atmosRepository.Repository{
		Config: r.Config,
	}
}
