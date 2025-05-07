package project

import (
	"context"

	"github.com/devpablocristo/monorepo/projects/qh/internal/project/handler/dto"
	"github.com/devpablocristo/monorepo/projects/qh/internal/project/usecases/domain"
)

type useCases struct {
	repo FullProjectsRepo
}

func NewUseCases(repo FullProjectsRepo) useCases {
	return useCases{
		repo: repo,
	}
}

func (uc *useCases) ProcessFullProjects(ctx context.Context, dtos []dto.FullProjectDTO) ([]string, error) {
	var domainProjects []domain.FullProject

	for _, dt := range dtos {
		d := dto.DtoToDomain(dt)

		domainProjects = append(domainProjects, d)
	}

	ids, err := uc.repo.SaveFullProject(ctx, domainProjects)
	if err != nil {
		return nil, err
	}

	return ids, nil
}
