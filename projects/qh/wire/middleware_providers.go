package wire

import (
	"github.com/gin-gonic/gin"

	mdw "github.com/devpablocristo/monorepo/pkg/rest/middlewares/gin"
)

func ProvideJwtMiddleware() (gin.HandlerFunc, error) {
	middleware := mdw.Validate(mdw.NewConfigFromEnv())
	return middleware, nil
}

func ProvideMiddlewares(jwtMiddleware gin.HandlerFunc) (*mdw.Middlewares, error) {
	globalMiddlewares := []gin.HandlerFunc{
		mdw.ErrorHandlingMiddleware(),
		mdw.RequestAndResponseLogger(mdw.HttpLoggingOptions{
			LogLevel:       "info",
			IncludeHeaders: true,
			IncludeBody:    false,
			ExcludedPaths: []string{
				"/health",
				"/ping",
				"/swagger/spec",
				"/swagger/ui/index.html",
			},
		}),
	}

	validatedMiddlewares := []gin.HandlerFunc{
		mdw.ValidateCredentials(),
	}

	protectedMiddlewares := []gin.HandlerFunc{
		jwtMiddleware,
	}

	return &mdw.Middlewares{
		Global:    globalMiddlewares,
		Validated: validatedMiddlewares,
		Protected: protectedMiddlewares,
	}, nil
}
