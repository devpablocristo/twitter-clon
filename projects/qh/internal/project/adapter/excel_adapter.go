package adapter

import (
	"fmt"
	"io"
	"strings"

	"github.com/devpablocristo/monorepo/projects/qh/internal/project/handler/dto"
	"github.com/xuri/excelize/v2"
)

type ExcelAdapter struct{}

func NewExcelAdapter() ExcelAdapter {
	return ExcelAdapter{}
}

func (a *ExcelAdapter) ParseFullProjectExcel(r io.Reader) ([]dto.FullProjectDTO, error) {
	f, err := excelize.OpenReader(r)
	if err != nil {
		return nil, fmt.Errorf("error opening excel: %w", err)
	}
	defer f.Close()

	var (
		projects []dto.FullProjectDTO
		errs     []string
	)

	for _, sheet := range f.GetSheetList() {
		rows, err := f.GetRows(sheet)
		if err != nil {
			errs = append(errs, fmt.Sprintf("failed to read sheet: %s: %v", sheet, err))
			continue
		}

		for i, row := range rows {
			if i == 0 {
				continue
			}

			proj, err := ParseRowToFullProject(row)
			if err != nil {
				errs = append(errs, fmt.Sprintf("row %d in sheet  '%s' : %v", i+1, sheet, err))
			}

			projects = append(projects, proj)
		}
	}

	if len(errs) > 0 {
		return projects, fmt.Errorf("errors while parsing file:\n%s", strings.Join(errs, "\n"))
	}

	return projects, nil
}
