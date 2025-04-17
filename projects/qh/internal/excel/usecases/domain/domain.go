package domain

import (
	"errors"

	"github.com/devpablocristo/monorepo/projects/qh/internal/excel/handler/dto"
)

type Person_2 struct {
	FirstName string
	LastName  string
	Age       int
	Phone     string
}

type Order struct {
	OrderID     string
	Customer    string
	TotalAmount float64
	Date        string
	Status      string
}

func NewOrder(dto dto.OrderDto) (*Order, error) {
	return &Order{
		OrderID:     dto.OrderID,
		Customer:    dto.Customer,
		TotalAmount: dto.TotalAmount,
		Date:        dto.Date,
		Status:      dto.Status,
	}, nil

}

func NewPerson(dto dto.ExcelPerson) (*Person_2, error) {
	if dto.Age < 0 || dto.Age > 120 {
		return nil, errors.New("invalid age")
	}
	if dto.FirstName == "" || dto.LastName == "" {
		return nil, errors.New("firstname & lastname required")
	}

	return &Person_2{
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Age:       dto.Age,
		Phone:     dto.Phone,
	}, nil
}
