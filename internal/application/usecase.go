package application

import (
	"github.com/umatare5/atmos-cli/internal/config"
	"github.com/umatare5/atmos-cli/internal/infrastructure"

	divelogUsecase "github.com/umatare5/atmos-cli/internal/application/divelog"
	userUsecase "github.com/umatare5/atmos-cli/internal/application/user"
)

// Usecase struct
type Usecase struct {
	Config     *config.Config
	Repository *infrastructure.Repository
}

// New returns Usecase struct
func New(c *config.Config, r *infrastructure.Repository) Usecase {
	return Usecase{
		Config:     c,
		Repository: r,
	}
}

// InvokeUserUsecase returns new Users Usecase
func (u *Usecase) InvokeUserUsecase() *userUsecase.Usecase {
	return &userUsecase.Usecase{
		Config:     u.Config,
		Repository: u.Repository,
	}
}

// InvokeDivelogUsecase returns new Divelog Usecase
func (u *Usecase) InvokeDivelogUsecase() *divelogUsecase.Usecase {
	return &divelogUsecase.Usecase{
		Config:     u.Config,
		Repository: u.Repository,
	}
}
