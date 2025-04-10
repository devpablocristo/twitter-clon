package excel

import (
	"fmt"
	"net/http"

	mdw "github.com/devpablocristo/monorepo/pkg/rest/middlewares/gin"
	gsv "github.com/devpablocristo/monorepo/pkg/rest/servers/gin"
	types "github.com/devpablocristo/monorepo/pkg/types"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	ucs useCases
	gsv gsv.Server
	mws *mdw.Middlewares
}

func NewHandler(s gsv.Server, u useCases, m *mdw.Middlewares) *Handler {
	return &Handler{
		ucs: u,
		gsv: s,
		mws: m,
	}
}

func (h *Handler) Routes() {
	router := h.gsv.GetRouter()

	apiVersion := h.gsv.GetApiVersion()
	apiBase := "/api" + apiVersion + "/persons"

	publicPrefix := apiBase + "public"
	validatePrefix := apiBase + "validated"
	protectedPrefix := apiBase + "/protected"

	public := router.Group(publicPrefix)
	{
		public.POST("/upload-excel", h.UploadExcel)
	}

	validated := router.Group(validatePrefix)
	{
		validated.Use(h.mws.Validated...)
	}

	protected := router.Group(protectedPrefix)
	{
		protected.Use(h.mws.Protected...)
		protected.GET("/ping", h.ProtectedPing)
	}
}

func (h *Handler) ProtectedPing(c *gin.Context) {
	c.JSON(http.StatusCreated, types.MessageResponse{
		Message: "Protected Pong!",
	})
}

func (h *Handler) UploadExcel(c *gin.Context) {
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		apiErr, code := types.NewAPIError(fmt.Errorf("no se pudo obtener el archivo %w", err))
		c.Error(apiErr).SetMeta(code)
		return
	}
	defer file.Close()

	if err := h.ucs.ProccesExcel(c.Request.Context(), file); err != nil {
		apiErr, code := types.NewAPIError(err)
		c.Error(apiErr).SetMeta(code)
		return
	}

	c.JSON(http.StatusCreated, types.MessageResponse{
		Message: "Archivo procesado correctamente",
	})
}
