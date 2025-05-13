package excel

import (
	"context"
	"fmt"

	"github.com/devpablocristo/monorepo/projects/qh/internal/excel/handler/dto"
	"github.com/devpablocristo/monorepo/projects/qh/internal/excel/usecases/domain"
)

type useCases struct {
	repoPerson Person_2Repository
	repoOrder  OrderRepository
}

func NewUseCases(r Person_2Repository, or OrderRepository) UseCases {
	return &useCases{
		repoPerson: r,
		repoOrder:  or,
	}
}

func (uc *useCases) ProccesPerson(ctx context.Context, dtos []dto.ExcelPerson) ([]string, error) {
	var errs []string
	var validPersons []domain.Person_2

	for i, dto := range dtos {
		person, err := domain.NewPerson(dto)
		if err != nil {
			errs = append(errs, fmt.Sprintf("fila %d, %v", i+1, err))
			continue
		}

		validPersons = append(validPersons, *person)
	}

	if len(validPersons) > 0 {
		if _, err := uc.repoPerson.SavePerson(ctx, validPersons); err != nil {
			return errs, err
		}
	}

	return errs, nil
}

func (uc *useCases) ProccesOrder(ctx context.Context, dtos []dto.OrderDto) ([]string, error) {
	var errs []string
	var validOrders []domain.Order

	for i, dto := range dtos {
		order, err := domain.NewOrder(dto)
		if err != nil {
			errs = append(errs, fmt.Sprintf("fila %d, %v", i+1, err))
			continue
		}

		validOrders = append(validOrders, *order)

	}
	if len(validOrders) > 0 {
		if _, err := uc.repoOrder.SaveOrder(ctx, validOrders); err != nil {
			return errs, err
		}
	}

	return errs, nil
}
