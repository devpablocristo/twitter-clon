package excel

import (
	"context"

	"github.com/devpablocristo/monorepo/projects/qh/internal/excel/handler/dto"
	"github.com/devpablocristo/monorepo/projects/qh/internal/excel/usecases/domain"
)

type Person_2Repository interface {
	SavePerson(context.Context, domain.Person_2) ([]string, error)
}

type UseCases interface {
	Procces(context.Context, []dto.ExcelPerson) ([]string, error)
}
