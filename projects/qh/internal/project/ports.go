package project

import (
	"context"

	"github.com/devpablocristo/monorepo/projects/qh/internal/project/handler/dto"
	"github.com/devpablocristo/monorepo/projects/qh/internal/project/usecases/domain"
)

type UseCases interface {
	ProcessFullProjects(context.Context, []dto.FullProjectDTO) ([]string, error)
}

type FullProjectsRepo interface {
	SaveFullProject(context.Context, []domain.FullProject) ([]string, error)
}
