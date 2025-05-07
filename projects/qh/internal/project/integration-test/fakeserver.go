package integrationtest_test

import (
	"context"
	"net/http"

	pkggin "github.com/devpablocristo/monorepo/pkg/rest/servers/gin"
	"github.com/gin-gonic/gin"
)

var _ pkggin.Server = (*fakeServer)(nil)

type fakeServer struct {
	engine     *gin.Engine
	apiVersion string
}

func newFakeServer(version string) pkggin.Server {
	return &fakeServer{
		engine:     gin.New(),
		apiVersion: version,
	}
}

func (s *fakeServer) GetRouter() *gin.Engine {
	return s.engine
}

func (s *fakeServer) GetApiVersion() string {
	return s.apiVersion
}

func (s *fakeServer) RunServer(ctx context.Context) error {
	return nil
}

func (s *fakeServer) WrapH(h http.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
