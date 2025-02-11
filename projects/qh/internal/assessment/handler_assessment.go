package assessment

import (
	"net/http"

	mdw "github.com/devpablocristo/monorepo/pkg/rest/middlewares/gin"
	types "github.com/devpablocristo/monorepo/pkg/types"
	utils "github.com/devpablocristo/monorepo/pkg/utils"
	dto "github.com/devpablocristo/monorepo/projects/qh/internal/assessment/handler/dto"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateAssessment(c *gin.Context) {
	userID, err := mdw.ExtractClaim(c, "sub", "")
	if err != nil {
		c.JSON(http.StatusUnauthorized, types.ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	var req dto.CreateAssessment
	if err := utils.ValidateRequest(c, &req); err != nil {
		apiErr, errCode := types.NewAPIError(err)
		c.Error(apiErr).SetMeta(errCode)
		return
	}
	req.HRID = userID

	ctx := c.Request.Context()
	assessment, err := req.ToDomain()
	if err != nil {
		apiErr, errCode := types.NewAPIError(err)
		c.Error(apiErr).SetMeta(errCode)
		return
	}

	newAssessmentID, err := h.ucs.CreateAssessment(ctx, assessment)
	if err != nil {
		apiErr, errCode := types.NewAPIError(err)
		c.Error(apiErr).SetMeta(errCode)
		return
	}

	c.JSON(http.StatusCreated, dto.CreateAssessmentResponse{
		Message:      "Assessment created successfully",
		AssessmentID: newAssessmentID,
	})
}

func (h *Handler) ListAssessments(c *gin.Context) {
	users, err := h.ucs.ListAssessments(c.Request.Context())
	if err != nil {
		apiErr, errCode := types.NewAPIError(err)
		c.Error(apiErr).SetMeta(errCode)
		return
	}
	c.JSON(http.StatusOK, users)
}

func (h *Handler) GetAssessment(c *gin.Context) {
	id := c.Param("id")

	assessment, err := h.ucs.GetAssessment(c.Request.Context(), id)
	if err != nil {
		apiErr, errCode := types.NewAPIError(err)
		c.Error(apiErr).SetMeta(errCode)
		return
	}

	c.JSON(http.StatusOK, assessment)
}

func (h *Handler) UpdateAssessment(c *gin.Context) {
	var req dto.Assessment
	if err := c.ShouldBindJSON(&req); err != nil {
		apiErr, errCode := types.NewAPIError(err)
		c.Error(apiErr).SetMeta(errCode)
		return
	}

	updatedAssessment, err := req.ToDomain()
	if err != nil {
		apiErr, errCode := types.NewAPIError(err)
		c.Error(apiErr).SetMeta(errCode)
		return
	}
	if err := h.ucs.UpdateAssessment(c.Request.Context(), updatedAssessment); err != nil {
		apiErr, errCode := types.NewAPIError(err)
		c.Error(apiErr).SetMeta(errCode)
		return
	}
	c.JSON(http.StatusCreated, types.MessageResponse{
		Message: "Assessment updated successfully",
	})
}

func (h *Handler) DeleteAssessment(c *gin.Context) {
	id := c.Param("id")
	if err := h.ucs.DeleteAssessment(c.Request.Context(), id); err != nil {
		apiErr, errCode := types.NewAPIError(err)
		c.Error(apiErr).SetMeta(errCode)
		return
	}
	c.JSON(http.StatusCreated, types.MessageResponse{
		Message: "Assessment deleted successfully",
	})
}
