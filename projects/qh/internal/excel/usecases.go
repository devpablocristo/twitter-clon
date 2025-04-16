package excel

import (
	"context"

	"github.com/devpablocristo/monorepo/projects/qh/internal/excel/handler/dto"
)

type useCases struct {
	repo Person_2Repository
}

func NewUseCases(r Person_2Repository) UseCases {
	return &useCases{repo: r}
}

func (uc *useCases) Procces(ctx context.Context, person []dto.ExcelPerson) ([]string, error) {
	return nil, nil
}
