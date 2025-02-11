package candidate

import (
	"net/http"

	"github.com/gin-gonic/gin"

	mdw "github.com/devpablocristo/monorepo/pkg/rest/middlewares/gin"
	gsv "github.com/devpablocristo/monorepo/pkg/rest/servers/gin"
	types "github.com/devpablocristo/monorepo/pkg/types"
	utils "github.com/devpablocristo/monorepo/pkg/utils"

	dto "github.com/devpablocristo/monorepo/projects/qh/internal/candidate/handler/dto"
)

type Handler struct {
	ucs UseCases
	gsv gsv.Server
	mws *mdw.Middlewares
}

func NewHandler(s gsv.Server, u UseCases, m *mdw.Middlewares) *Handler {
	return &Handler{
		ucs: u,
		gsv: s,
		mws: m,
	}
}

func (h *Handler) Routes() {
	router := h.gsv.GetRouter()

	apiVersion := h.gsv.GetApiVersion()
	apiBase := "/api/" + apiVersion + "/candidates"
	//publicPrefix := apiBase + "/public"
	validatedPrefix := apiBase + "/validated"
	protectedPrefix := apiBase + "/protected"

	// Rutas públicas
	// public := router.Group(publicPrefix)
	// {
	// 	public.POST("", h.CreateCandidate)
	// 	public.GET("", h.ListCandidates)
	// 	public.GET("/:id", h.GetCandidate)
	// 	public.PUT("/:id", h.UpdateCandidate)
	// 	public.DELETE("/:id", h.DeleteCandidate)
	// }

	validated := router.Group(validatedPrefix)
	{
		// Aplicar middleware de validación de credenciales
		validated.Use(h.mws.Validated...)
		// Puedes añadir rutas aquí si es necesario
	}

	// Rutas protegidas
	protected := router.Group(protectedPrefix)
	{
		protected.Use(h.mws.Protected...)

		protected.GET("/ping", h.ProtectedPing)

		protected.POST("", h.CreateCandidate)
		protected.GET("", h.ListCandidates)
		protected.GET("/:id", h.GetCandidate)
		protected.PUT("/:id", h.UpdateCandidate)
		protected.DELETE("/:id", h.DeleteCandidate)
	}
}

func (h *Handler) ProtectedPing(c *gin.Context) {
	c.JSON(http.StatusCreated, types.MessageResponse{
		Message: "Protected Pong!",
	})
}

func (h *Handler) CreateCandidate(c *gin.Context) {
	var req dto.CreateCandidate
	if err := utils.ValidateRequest(c, &req); err != nil {
		apiErr, errCode := types.NewAPIError(err)
		c.Error(apiErr).SetMeta(errCode)
		return
	}

	ctx := c.Request.Context()
	newCandidateID, err := h.ucs.CreateCandidate(ctx, req.ToDomain())
	if err != nil {
		apiErr, errCode := types.NewAPIError(err)
		c.Error(apiErr).SetMeta(errCode)
		return
	}

	c.JSON(http.StatusCreated, dto.CreateCandidateResponse{
		Message:     "Candidate created successfully",
		CandidateID: newCandidateID,
	})
}

func (h *Handler) ListCandidates(c *gin.Context) {
	candidates, err := h.ucs.ListCandidates(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error fetching candidates"})
		return
	}
	c.JSON(http.StatusOK, candidates)
}

func (h *Handler) GetCandidate(c *gin.Context) {
	id := c.Param("id")

	person, err := h.ucs.GetCandidate(c.Request.Context(), id)
	if err != nil {
		apiErr, errCode := types.NewAPIError(err)
		c.Error(apiErr).SetMeta(errCode)
		return
	}

	c.JSON(http.StatusOK, person)
}

func (h *Handler) UpdateCandidate(c *gin.Context) {
	var updatedCandidate dto.Candidate
	if err := c.ShouldBindJSON(&updatedCandidate); err != nil {
		apiErr, errCode := types.NewAPIError(err)
		c.Error(apiErr).SetMeta(errCode)
		return
	}

	if err := h.ucs.UpdateCandidate(c.Request.Context(), updatedCandidate.ToDomain()); err != nil {
		apiErr, errCode := types.NewAPIError(err)
		c.Error(apiErr).SetMeta(errCode)
		return
	}
	c.JSON(http.StatusCreated, types.MessageResponse{
		Message: "Candidate updated successfully",
	})
}

func (h *Handler) DeleteCandidate(c *gin.Context) {
	id := c.Param("id")
	if err := h.ucs.DeleteCandidate(c.Request.Context(), id); err != nil {
		apiErr, errCode := types.NewAPIError(err)
		c.Error(apiErr).SetMeta(errCode)
		return
	}
	c.JSON(http.StatusCreated, types.MessageResponse{
		Message: "Candidate deleted successfully",
	})
}
