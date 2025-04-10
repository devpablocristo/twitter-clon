package excel

import (
	"context"
	"mime/multipart"

	"github.com/devpablocristo/monorepo/projects/qh/internal/excel/usecases/domain"
)

type Person_2Repository interface {
	CreatePerson(context.Context, domain.Person_2) error
}

type UseCases interface {
	ProccesExcel(context.Context, multipart.File) error
}
