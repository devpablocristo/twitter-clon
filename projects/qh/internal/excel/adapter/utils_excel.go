package adapter

import (
	"fmt"
	"strconv"

	"github.com/devpablocristo/monorepo/projects/qh/internal/excel/handler/dto"
)

func parseRowToPerson(row []string) (dto.ExcelPerson, error) {
	if len(row) < 4 {
		return dto.ExcelPerson{}, fmt.Errorf("incomplete data")
	}

	age, err := strconv.Atoi(row[2])
	if err != nil {
		return dto.ExcelPerson{}, fmt.Errorf("invalid age: %v", err)
	}

	return dto.ExcelPerson{
		FirstName: row[0],
		LastName:  row[1],
		Age:       age,
		Phone:     row[3],
	}, nil
}

func parseRowToOrder(row []string) (dto.OrderDto, error) {
	if len(row) < 5 {
		return dto.OrderDto{}, fmt.Errorf("imcomplete data")
	}

	total, err := strconv.ParseFloat(row[2], 64)
	if err != nil {
		return dto.OrderDto{}, fmt.Errorf("invalid amount: %v", err)
	}

	return dto.OrderDto{
		OrderID:     row[0],
		Customer:    row[1],
		TotalAmount: total,
		Status:      row[3],
		Date:        row[4],
	}, nil
}
