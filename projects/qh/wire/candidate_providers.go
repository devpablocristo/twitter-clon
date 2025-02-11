package wire

import (
	"errors"

	gorm "github.com/devpablocristo/monorepo/pkg/databases/sql/gorm"
	mdw "github.com/devpablocristo/monorepo/pkg/rest/middlewares/gin"
	ginsrv "github.com/devpablocristo/monorepo/pkg/rest/servers/gin"

	candidate "github.com/devpablocristo/monorepo/projects/qh/internal/candidate"
)

func ProvideCandidateRepository(repo gorm.Repository) (candidate.Repository, error) {
	if repo == nil {
		return nil, errors.New("gorm repository cannot be nil")
	}
	return candidate.NewRepository(repo), nil
}

func ProvideCandidateUseCases(repo candidate.Repository) candidate.UseCases {
	return candidate.NewUseCases(repo)
}

func ProvideCandidateHandler(server ginsrv.Server, usecases candidate.UseCases, middlewares *mdw.Middlewares) *candidate.Handler {
	return candidate.NewHandler(server, usecases, middlewares)
}
