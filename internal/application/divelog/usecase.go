package usecase

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/umatare5/atmos-cli/internal/config"
	"github.com/umatare5/atmos-cli/internal/domain"
	"github.com/umatare5/atmos-cli/internal/infrastructure"

	"github.com/umatare5/atmos-go"
)

// Usecase struct
type Usecase struct {
	Config     *config.Config
	Repository *infrastructure.Repository
}

// GetDivelogs ...
func (u *Usecase) GetDivelogs() (*[]byte, error) {
	data, err := u.Repository.InvokeAtmosRepository().GetDivelogs()
	if err != nil {
		return nil, err
	}

	json, err := json.Marshal(data.JSON200)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &json, nil
}

// GetDivelogsPretty returns divelogs ...
func (u *Usecase) GetDivelogsPretty() (*table.Row, *[]table.Row, error) {
	data, err := u.Repository.InvokeAtmosRepository().GetDivelogs()
	if err != nil {
		return nil, nil, err
	}

	return &table.Row{"Region", "Point", "Entry Time", "Dive Time", "Max Depth", "Divelog ID"},
		u.minimizeDivelogs(data),
		nil
}

// GetDivelog ...
func (u *Usecase) GetDivelog(divelogID string) (*atmos.GetDivelogResponse, error) {
	data, err := u.Repository.InvokeAtmosRepository().GetDivelog(divelogID)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (u *Usecase) minimizeDivelogs(data *atmos.GetDivelogsResponse) *[]table.Row {
	var divelogs []table.Row
	for _, v := range *data.JSON200.Response.Divelogs {
		minutes := strconv.Itoa(*v.DiveDuration / 60)
		seconds := strconv.Itoa(*v.DiveDuration % 60)

		poiName := domain.NotAvailableString
		if v.Poi != nil {
			poiName = *v.Poi.PoiName
		}

		poiRegion := domain.NotAvailableString
		if v.Poi != nil {
			poiRegion = *v.Poi.PoiRegion
		}

		divelogs = append(divelogs, table.Row{
			poiRegion,
			poiName,
			time.Unix(int64(*v.DiveDatetime), 0),
			minutes + "m " + seconds + "s",
			*v.MaxDepth + "m",
			*v.DivelogId,
		})
	}

	return &divelogs
}
