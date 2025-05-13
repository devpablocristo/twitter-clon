package excel

import (
	"context"

	"github.com/devpablocristo/monorepo/projects/qh/internal/excel/handler/dto"
	"github.com/devpablocristo/monorepo/projects/qh/internal/excel/usecases/domain"
)

type Person_2Repository interface {
	SavePerson(context.Context, []domain.Person_2) ([]string, error)
}

type OrderRepository interface {
	SaveOrder(context.Context, []domain.Order) ([]string, error)
}

type UseCases interface {
	ProccesPerson(context.Context, []dto.ExcelPerson) ([]string, error)
	ProccesOrder(context.Context, []dto.OrderDto) ([]string, error)
}
