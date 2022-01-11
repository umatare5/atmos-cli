package framework

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/umatare5/atmos-cli/internal/application"
	"github.com/umatare5/atmos-cli/internal/config"
	"github.com/umatare5/atmos-cli/internal/infrastructure"
)

// Exec struct
type Exec struct {
	Config     *config.Config
	Repository *infrastructure.Repository
	Usecase    *application.Usecase
}

// New returns Exec struct
func New(c *config.Config, r *infrastructure.Repository, u *application.Usecase) Exec {
	return Exec{
		Config:     c,
		Repository: r,
		Usecase:    u,
	}
}

// GetUserProfile ...
func (e *Exec) GetUserProfile(userID string, PrettyFormat bool) {
	data, err := e.Usecase.InvokeUserUsecase().GetUserProfile(userID)
	if err != nil {
		fmt.Println(err)
		return
	}

	json, err := json.Marshal(data.JSON200)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(json))
}

// GetUserStatistics ...
func (e *Exec) GetUserStatistics(userID string, PrettyFormat bool) {
	data, err := e.Usecase.InvokeUserUsecase().GetUserStatistics(userID)
	if err != nil {
		fmt.Println(err)
		return
	}

	json, err := json.Marshal(data.JSON200)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(json))
}

// GetDivelogs ...
func (e *Exec) GetDivelogs() {
	data, err := e.Usecase.InvokeDivelogUsecase().GetDivelogs()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(*data))
}

// GetDivelogsPretty ...
func (e *Exec) GetDivelogsPretty() {
	header, rows, err := e.Usecase.InvokeDivelogUsecase().GetDivelogsPretty()
	if err != nil {
		fmt.Println(err)
		return
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(*header)
	t.AppendRows(*rows)
	t.AppendSeparator()
	t.Render()
}

// GetDivelog ...
func (e *Exec) GetDivelog(divelogID string, PrettyFormat bool) {
	data, err := e.Usecase.InvokeDivelogUsecase().GetDivelog(divelogID)
	if err != nil {
		fmt.Println(err)
		return
	}

	json, err := json.Marshal(data.JSON200)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(json))
}
