package adapter

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/devpablocristo/monorepo/projects/qh/internal/project/handler/dto"
)

func ParseRowToFullProject(row []string) (dto.FullProjectDTO, error) {
	var dto dto.FullProjectDTO
	var errMsg []string

	if len(row) < 16 {
		return dto, fmt.Errorf("expected at least 16 columns, got %d", len(row))
	}

	dto.ClientName = row[0]
	dto.ProjectName = row[1]
	dto.ManagerName = row[2]
	dto.ManagerEmail = row[3]
	dto.ManagerTitle = row[4]
	dto.InvestorName = row[5]
	dto.InvestorEmail = row[6]
	dto.InvestorCompany = row[7]

	dto.InvestmentPct, _ = strconv.ParseFloat(row[8], 64)
	dto.FieldName = row[9]
	dto.LeaseType = row[10]
	dto.PlotName = row[11]
	dto.Hectares, _ = strconv.ParseFloat(row[12], 64)
	dto.PreviousCrop = row[13]
	dto.CurrentCrop = row[14]
	dto.Season = row[15]

	if len(errMsg) > 0 {
		return dto, fmt.Errorf(strings.Join(errMsg, ";"))
	}

	return dto, nil
}
