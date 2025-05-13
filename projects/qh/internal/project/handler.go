package project

import (
	"fmt"
	"net/http"

	mdw "github.com/devpablocristo/monorepo/pkg/rest/middlewares/gin"
	gsv "github.com/devpablocristo/monorepo/pkg/rest/servers/gin"
	"github.com/devpablocristo/monorepo/projects/qh/internal/project/adapter"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	usc useCases
	gsv gsv.Server
	mdw *mdw.Middlewares
	ea  adapter.ExcelAdapter
}

func NewHandler(s gsv.Server, u useCases, m *mdw.Middlewares, e adapter.ExcelAdapter) *Handler {
	return &Handler{
		usc: u,
		gsv: s,
		mdw: m,
		ea:  e,
	}
}

func (h *Handler) Routes() {
	router := h.gsv.GetRouter()
	apiVersion := h.gsv.GetApiVersion()
	apiBase := "/api/" + apiVersion + "/upload"

	fullProject := router.Group(apiBase + "/fullProject")
	{
		fullProject.POST("", h.UploadFullProjectExcel)
	}
}

func (h *Handler) UploadFullProjectExcel(c *gin.Context) {

	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("could not read file: %v", err)})
		return
	}

	// parseo el excel a DTOs
	fullProjects, err := h.ea.ParseFullProjectExcel(file)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "error parsing file",
			"details": err.Error(),
		})
		return
	}

	if len(fullProjects) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no valid projects found in file"})
		return
	}

	// procesando los projects
	ids, err := h.usc.ProcessFullProjects(c.Request.Context(), fullProjects)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "failed to save projects",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":        "projects processed successfully",
		"project_ids":    ids,
		"projects_count": len(ids),
	})
}
