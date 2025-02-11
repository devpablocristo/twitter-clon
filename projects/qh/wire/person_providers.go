package wire

import (
	"errors"

	pgdb "github.com/devpablocristo/monorepo/pkg/databases/sql/postgresql/pgxpool"
	mdw "github.com/devpablocristo/monorepo/pkg/rest/middlewares/gin"
	ginsrv "github.com/devpablocristo/monorepo/pkg/rest/servers/gin"

	person "github.com/devpablocristo/monorepo/projects/qh/internal/person"
)

func ProvidePersonRepository(repo pgdb.Repository) (person.Repository, error) {
	if repo == nil {
		return nil, errors.New("postgres repository cannot be nil")
	}
	return person.NewPostgresRepository(repo), nil
}

func ProvidePersonUseCases(repo person.Repository) person.UseCases {
	return person.NewUseCases(repo)
}

func ProvidePersonHandler(server ginsrv.Server, usecases person.UseCases, middlewares *mdw.Middlewares) *person.Handler {
	return person.NewHandler(server, usecases, middlewares)
}
