// wire.go
package wire

import (
	"errors"

	jwt "github.com/devpablocristo/monorepo/pkg/authe/jwt/v5"
	redis "github.com/devpablocristo/monorepo/pkg/databases/cache/redis/v8"
	resty "github.com/devpablocristo/monorepo/pkg/rest/clients/resty"
	mdw "github.com/devpablocristo/monorepo/pkg/rest/middlewares/gin"
	ginsrv "github.com/devpablocristo/monorepo/pkg/rest/servers/gin"

	authe "github.com/devpablocristo/monorepo/projects/qh/internal/authe"
	config "github.com/devpablocristo/monorepo/projects/qh/internal/config"
)

// ProvideAutheCache proporciona una implementación de authe.Cache utilizando Redis.
func ProvideAutheCache(cache redis.Cache) (authe.Cache, error) {
	if cache == nil {
		return nil, errors.New("redis cache cannot be nil")
	}
	return authe.NewRedisCache(cache), nil
}

// ProvideAutheJwtService proporciona una implementación de authe.JwtService utilizando el servicio JWT.
func ProvideAutheJwtService(jwtSrv jwt.Service, cnfLdr config.Loader) (authe.JwtService, error) {
	if jwtSrv == nil {
		return nil, errors.New("jwt service cannot be nil")
	}
	return authe.NewJwtService(jwtSrv, cnfLdr)
}

// ProvideAutheHttpClient proporciona una implementación de authe.HttpClient utilizando Resty.
func ProvideAutheHttpClient(httpc resty.Client, cnfLdr config.Loader) (authe.HttpClient, error) {
	if httpc == nil {
		return nil, errors.New("http client cannot be nil")
	}
	return authe.NewHttpClient(httpc, cnfLdr), nil
}

// ProvideAutheUseCases proporciona una implementación de authe.UseCases con todas sus dependencias.
func ProvideAutheUseCases(ch authe.Cache, js authe.JwtService, hc authe.HttpClient) authe.UseCases {
	return authe.NewUseCases(ch, js, hc)
}

// ProvideAutheHandler proporciona un controlador de authe.Handler configurado con el servidor, casos de uso y middlewares.
func ProvideAutheHandler(server ginsrv.Server, usecases authe.UseCases, middlewares *mdw.Middlewares) *authe.Handler {
	return authe.NewHandler(server, usecases, middlewares)
}
