package wire

import (
	mdw "github.com/devpablocristo/monorepo/pkg/rest/middlewares/gin"
	ginsrv "github.com/devpablocristo/monorepo/pkg/rest/servers/gin"
	"github.com/devpablocristo/monorepo/projects/qh/internal/excel"
	"github.com/devpablocristo/monorepo/projects/qh/internal/excel/adapter"
	"go.mongodb.org/mongo-driver/mongo"
)

func ProvideMongoRepository(db *mongo.Database) *excel.MongoRepository {
	return excel.NewMongoRepository(db)
}

func ProvidePerson_2Usecase(r excel.Person_2Repository) excel.UseCases {
	return excel.NewUseCases(r)
}

func ProvideExcelHandler(server ginsrv.Server, usecases excel.UseCases, middleware *mdw.Middlewares, excelAdapter adapter.ExcelAdapter) *excel.Handler {
	return excel.NewHandler(server, usecases, middleware, excelAdapter)
}
