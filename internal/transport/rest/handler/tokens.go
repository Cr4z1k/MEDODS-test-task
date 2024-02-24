package handler

import (
	"net/http"

	"github.com/Cr4z1k/MEDODS-test-task/internal/core"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetTokens(c *gin.Context) {
	guid := c.Param("guid")

	tokens, err := h.s.GetTokens(guid)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tokens)
}

func (h *Handler) RefreshAccess(c *gin.Context) {
	var refresh core.Refresh

	if err := c.ShouldBindJSON(&refresh); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	tokens, err := h.s.RefreshTokens(refresh.RefreshToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tokens)
}
