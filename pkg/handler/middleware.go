package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponce(c, http.StatusUnauthorized, "empty header")
		return
	}
	headerParts := strings.Split(header, "")
	if len(headerParts) != 2 {
		newErrorResponce(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	userId, err := h.services.Authorization.ParseTokern(headerParts[1])
	if err != nil {
		newErrorResponce(c, http.StatusUnauthorized, err.Error())
		return
	}
	c.Set("userCtx", userId)
}

// func getUderId(c *gin.Context) (int, error) {
// 	id, ok := c.Get(userCtx)
// 	if !ok {
// 		newErrorResponce(c, http.StatusInternalServerError, "user id is not foud")
// 		return 0, errors.New("user is not found")
// 	}

// 	idInt, ok := id.(int)
// 	if !ok {
// 		newErrorResponce(c, http.StatusInternalServerError, "user id is not foud")
// 		return 0, errors.New("user is not found")
// 	}
// 	return idInt, nil
// }
