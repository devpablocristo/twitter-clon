package excel

import (
	"fmt"
	"log"
	"net/http"

	mdw "github.com/devpablocristo/monorepo/pkg/rest/middlewares/gin"
	gsv "github.com/devpablocristo/monorepo/pkg/rest/servers/gin"
	types "github.com/devpablocristo/monorepo/pkg/types"
	"github.com/devpablocristo/monorepo/projects/qh/internal/excel/adapter"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	ucs UseCases
	gsv gsv.Server
	mws *mdw.Middlewares
	ea  adapter.ExcelAdapter
}

func NewHandler(s gsv.Server, u UseCases, m *mdw.Middlewares, e adapter.ExcelAdapter) *Handler {
	return &Handler{
		ucs: u,
		gsv: s,
		mws: m,
		ea:  e,
	}
}

func (h *Handler) Routes() {
	router := h.gsv.GetRouter()

	apiVersion := h.gsv.GetApiVersion()
	apiBase := "/api" + apiVersion + "/upload"

	publicPrefix := apiBase + "public"
	validatePrefix := apiBase + "validated"
	protectedPrefix := apiBase + "/protected"

	public := router.Group(publicPrefix)
	{
		public.POST("/person", h.UploadPersonExcel)
		public.POST("/order", h.UploadOrderExcel)
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

func (h *Handler) UploadPersonExcel(c *gin.Context) {
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		apiErr, code := types.NewAPIError(fmt.Errorf("no se pudo obtener el archivo %w", err))
		c.Error(apiErr).SetMeta(code)
		return
	}
	defer file.Close()

	// llamo al adaptador excel
	persons, err := h.ea.ParsePersonExcel(file)
	if err != nil {
		log.Fatalf("error parsing excel: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "internal error"})
		return
	}

	// llamo al caso de uso
	if len(persons) > 0 {
		if errList, err := h.ucs.ProccesPerson(c.Request.Context(), persons); err != nil {
			log.Printf("Error processing data: %v", errList)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "error processing data",
				"details": errList,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "file processed successfully"})
	}

}

func (h *Handler) UploadOrderExcel(c *gin.Context) {
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		apiErr, code := types.NewAPIError(fmt.Errorf("no se pudo obtener el archivo %w", err))
		c.Error(apiErr).SetMeta(code)
		return
	}
	defer file.Close()

	orders, err := h.ea.ParseOrderExcel(file)
	if err != nil {
		log.Fatalf("error parsing excel: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "internal error"})
		return
	}

	if len(orders) > 0 {
		if errList, err := h.ucs.ProccesOrder(c.Request.Context(), orders); err != nil {
			log.Printf("Error processing data: %v", errList)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "error processing data",
				"details": errList,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "file processed successfully"})
	}

}
