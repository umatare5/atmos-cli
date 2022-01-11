package repository

import (
	"context"

	"github.com/umatare5/atmos-cli/internal/config"
	client "github.com/umatare5/atmos-cli/pkg/atmos"

	"github.com/umatare5/atmos-go"
)

// Repository struct
type Repository struct {
	Config *config.Config
}

// GetUser ...
func (r *Repository) GetUser(userID string) (*atmos.GetUserResponse, error) {
	client, err := client.NewClientWithToken(r.Config.Server, r.Config.AccessToken)
	if err != nil {
		return nil, err
	}

	data, err := client.GetUserWithResponse(
		context.Background(), userID,
	)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// GetUserStatistics ...
func (r *Repository) GetUserStatistics(userID string) (*atmos.GetUserStatisticsResponse, error) {
	client, err := client.NewClientWithToken(r.Config.Server, r.Config.AccessToken)
	if err != nil {
		return nil, err
	}

	data, err := client.GetUserStatisticsWithResponse(
		context.Background(), userID,
	)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// GetDivelogs ...
func (r *Repository) GetDivelogs() (*atmos.GetDivelogsResponse, error) {
	client, err := client.NewClientWithToken(r.Config.Server, r.Config.AccessToken)
	if err != nil {
		return nil, err
	}

	data, err := client.GetDivelogsWithResponse(
		context.Background(),
		&atmos.GetDivelogsParams{
			Limit:  &r.Config.Limit,
			Cursor: &r.Config.Cursor,
		},
	)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// GetDeletedDivelogs ...
func (r *Repository) GetDeletedDivelogs() (*atmos.GetDeletedDivelogsResponse, error) {
	client, err := client.NewClientWithToken(r.Config.Server, r.Config.AccessToken)
	if err != nil {
		return nil, err
	}

	data, err := client.GetDeletedDivelogsWithResponse(
		context.Background(),
		&atmos.GetDeletedDivelogsParams{
			Limit:  &r.Config.Limit,
			Cursor: &r.Config.Cursor,
		},
	)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// GetDivelog ...
func (r *Repository) GetDivelog(divelogID string) (*atmos.GetDivelogResponse, error) {
	client, err := client.NewClientWithToken(r.Config.Server, r.Config.AccessToken)
	if err != nil {
		return nil, err
	}

	data, err := client.GetDivelogWithResponse(
		context.Background(), divelogID,
	)
	if err != nil {
		return nil, err
	}

	return data, nil
}
