package adapter

import (
	"fmt"
	"io"
	"strings"

	"github.com/devpablocristo/monorepo/projects/qh/internal/excel/handler/dto"
	"github.com/xuri/excelize/v2"
)

type ExcelAdapter struct{}

func NewExcelAdapter() ExcelAdapter {
	return ExcelAdapter{}
}

func (a *ExcelAdapter) ParsePersonExcel(r io.Reader) ([]dto.ExcelPerson, error) {
	f, err := excelize.OpenReader(r)
	if err != nil {
		return nil, fmt.Errorf("error opening excel %w", err)
	}
	defer f.Close()

	var (
		persons []dto.ExcelPerson
		errs    []string
	)

	for _, sheet := range f.GetSheetList() {
		rows, err := f.GetRows(sheet)
		if err != nil {
			errs = append(errs, fmt.Sprintf("failed to read sheet %s: %v", sheet, err))
		}

		for i, row := range rows {
			if i == 0 {
				continue
			}

			person, err := parseRowToPerson(row)
			if err != nil {
				errs = append(errs, fmt.Sprintf("row %d in sheet '%s' : %v", i+1, sheet, err))
			}

			persons = append(persons, person)
		}
	}

	if len(errs) > 0 {
		return persons, fmt.Errorf("errors while parsing file:\n%s", strings.Join(errs, "\n"))
	}

	return persons, nil
}

func (a *ExcelAdapter) ParseOrderExcel(r io.Reader) ([]dto.OrderDto, error) {
	f, err := excelize.OpenReader(r)
	if err != nil {
		return nil, fmt.Errorf("error opening excel: %w", err)
	}
	defer f.Close()

	var (
		orders []dto.OrderDto
		errs   []string
	)

	for _, sheet := range f.GetSheetList() {
		rows, err := f.GetRows(sheet)
		if err != nil {
			errs = append(errs, fmt.Sprintf("failed to read sheet %s: %v", sheet, err))
			continue
		}

		for i, row := range rows {

			if i == 0 {
				continue
			}

			order, err := parseRowToOrder(row)
			if err != nil {
				errs = append(errs, fmt.Sprintf("row %d in sheet '%s': %v", i+1, sheet, err))
			}

			orders = append(orders, order)
		}
	}

	if len(errs) > 0 {
		return orders, fmt.Errorf("errors while parsing file:\n%s", strings.Join(errs, "\n"))
	}

	return orders, nil
}
