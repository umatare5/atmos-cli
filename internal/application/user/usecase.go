package usecase

import (
	"github.com/umatare5/atmos-cli/internal/config"
	"github.com/umatare5/atmos-cli/internal/infrastructure"

	"github.com/umatare5/atmos-go"
)

// Usecase struct
type Usecase struct {
	Config     *config.Config
	Repository *infrastructure.Repository
}

// GetUserProfile ...
func (u *Usecase) GetUserProfile(userID string) (*atmos.GetUserResponse, error) {
	data, err := u.Repository.InvokeAtmosRepository().GetUser(userID)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// GetUserStatistics ...
func (u *Usecase) GetUserStatistics(userID string) (*atmos.GetUserStatisticsResponse, error) {
	data, err := u.Repository.InvokeAtmosRepository().GetUserStatistics(userID)
	if err != nil {
		return nil, err
	}
	return data, nil
}
