package handler

import (
	"net/http"

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

}
